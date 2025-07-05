package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func main() {
	var err error
	db, err = setupDatabase()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
	defer db.Close()

	if err := createTableIfNotExists(); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	http.HandleFunc("/process", processHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("TX service starting on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func setupDatabase() (*pgxpool.Pool, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:postgres@postgres:5432/testdb?sslmode=disable"
	}

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	config.MaxConns = 30
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = time.Minute * 30

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	return pool, nil
}

func createTableIfNotExists() error {
	query := `
	CREATE TABLE IF NOT EXISTS requests (
		id SERIAL PRIMARY KEY,
		request_id VARCHAR(255) NOT NULL,
		processed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		external_call_duration INTEGER
	);`

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	log.Println("TX service: requests table ready")
	return nil
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	requestID := generateRequestID()
	startTime := time.Now()
	log.Printf("TX service: starting request %s", requestID)

	stats := db.Stat()
	log.Printf("TX service: connection pool stats - acquired: %d, idle: %d, max: %d", 
		stats.AcquiredConns(), stats.IdleConns(), stats.MaxConns())

	tx, err := db.Begin(context.Background())
	if err != nil {
		log.Printf("TX service: failed to begin transaction for request %s: %v", requestID, err)
		http.Error(w, "Connection pool exhausted", http.StatusServiceUnavailable)
		return
	}
	defer tx.Rollback(context.Background())

	log.Printf("TX service: transaction started for request %s", requestID)

	slowResponse, err := callSlowService(requestID)
	if err != nil {
		log.Printf("TX service: slow service call failed for request %s: %v", requestID, err)
		http.Error(w, "External service error", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec(context.Background(),
		"INSERT INTO requests (request_id, external_call_duration) VALUES ($1, $2)",
		requestID, 5000)
	if err != nil {
		log.Printf("TX service: insert failed for request %s: %v", requestID, err)
		http.Error(w, "Database insert failed", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(context.Background()); err != nil {
		log.Printf("TX service: commit failed for request %s: %v", requestID, err)
		http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
		return
	}

	duration := time.Since(startTime)
	log.Printf("TX service: completed request %s in %v", requestID, duration)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"request_id": requestID,
		"duration_ms": duration.Milliseconds(),
		"slow_service_response": slowResponse,
	})
}

func callSlowService(requestID string) (map[string]interface{}, error) {
	slowServiceURL := os.Getenv("SLOW_SERVICE_URL")
	if slowServiceURL == "" {
		slowServiceURL = "http://slow-service:8080"
	}

	log.Printf("TX service: calling slow service for request %s", requestID)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(slowServiceURL + "/slow")
	if err != nil {
		return nil, fmt.Errorf("slow service call failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("slow service returned %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	log.Printf("TX service: slow service call completed for request %s", requestID)
	return result, nil
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	stats := db.Stat()
	health := map[string]interface{}{
		"status": "healthy",
		"connection_pool": map[string]interface{}{
			"acquired_conns": stats.AcquiredConns(),
			"idle_conns":     stats.IdleConns(),
			"max_conns":      stats.MaxConns(),
			"total_conns":    stats.TotalConns(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

func generateRequestID() string {
	return uuid.New().String()[:8]
}