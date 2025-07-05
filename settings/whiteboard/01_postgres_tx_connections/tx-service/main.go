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

	server := &http.Server{
		Addr: ":8081",
	}

	log.Println("TX service starting on :8081")
	if err := server.ListenAndServe(); err != nil {
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

	// Configure connection pool to handle high concurrency
	config.MaxConns = 100  // Match PostgreSQL limit
	config.MinConns = 10   // Keep some ready

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

	tx, err := db.Begin(context.Background())
	if err != nil {
		http.Error(w, "Connection pool exhausted", http.StatusServiceUnavailable)
		return
	}
	defer tx.Rollback(context.Background())

	slowResponse, err := callSlowService()
	if err != nil {
		http.Error(w, "External service error", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec(
		context.Background(),
		"INSERT INTO requests (request_id, external_call_duration) VALUES ($1, $2)",
		requestID, 5000,
	)
	if err != nil {
		http.Error(w, "Database insert failed", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(context.Background()); err != nil {
		http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
		return
	}

	duration := time.Since(startTime)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"request_id":            requestID,
			"duration_ms":           duration.Milliseconds(),
			"slow_service_response": slowResponse,
		},
	)
}

func callSlowService() (map[string]interface{}, error) {
	slowServiceURL := "http://slow-service:8080"
	client := &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", slowServiceURL+"/slow", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
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
