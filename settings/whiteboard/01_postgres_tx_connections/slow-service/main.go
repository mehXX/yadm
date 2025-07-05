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
	totalRequests  int64
	activeRequests int64
	requestCounter int64
)

func main() {
	http.HandleFunc("/slow", slowHandler)

	log.Printf("Slow service starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	requestID := generateRequestID()
	atomic.AddInt64(&totalRequests, 1)
	atomic.AddInt64(&activeRequests, 1)

	time.Sleep(3 * time.Second)

	response := map[string]interface{}{
		"request_id": requestID,
		"timestamp":  time.Now(),
		"message":    "Response after 3 seconds",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	atomic.AddInt64(&activeRequests, -1)
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
