package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	totalRequests int64
	activeRequests int64
	requestCounter int64
)

func main() {
	go logMetrics()
	
	http.HandleFunc("/slow", slowHandler)
	
	log.Printf("Slow service starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	requestID := generateRequestID()
	atomic.AddInt64(&totalRequests, 1)
	atomic.AddInt64(&activeRequests, 1)
	
	log.Printf("Slow service: received request %s", requestID)
	
	time.Sleep(5 * time.Second)
	
	response := map[string]interface{}{
		"request_id": requestID,
		"timestamp": time.Now(),
		"message": "Response after 5 seconds",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	
	atomic.AddInt64(&activeRequests, -1)
	log.Printf("Slow service: completed request %s", requestID)
}

func generateRequestID() string {
	counter := atomic.AddInt64(&requestCounter, 1)
	return fmt.Sprintf("req_%d_%s", counter, randomString(6))
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func logMetrics() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	
	var lastTotal int64
	
	for {
		select {
		case <-ticker.C:
			currentTotal := atomic.LoadInt64(&totalRequests)
			currentActive := atomic.LoadInt64(&activeRequests)
			
			rps := float64(currentTotal-lastTotal) / 1.0
			lastTotal = currentTotal
			
			log.Printf("Metrics: total=%d, active=%d, rps=%.2f", currentTotal, currentActive, rps)
		}
	}
}