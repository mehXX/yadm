package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type TestResult struct {
	RequestNum int           `json:"request_num"`
	StartTime  time.Time     `json:"start_time"`
	EndTime    time.Time     `json:"end_time"`
	Duration   time.Duration `json:"duration_ms"`
	StatusCode int           `json:"status_code"`
	Success    bool          `json:"success"`
	Error      string        `json:"error,omitempty"`
}

type TestSummary struct {
	TotalRequests  int            `json:"total_requests"`
	SuccessCount   int            `json:"success_count"`
	FailureCount   int            `json:"failure_count"`
	SuccessRate    float64        `json:"success_rate"`
	TotalDuration  float64        `json:"total_duration_seconds"`
	AverageLatency float64        `json:"average_latency_seconds"`
	MinLatency     float64        `json:"min_latency_seconds"`
	MaxLatency     float64        `json:"max_latency_seconds"`
	P95Latency     float64        `json:"p95_latency_seconds"`
	ErrorsByType   map[string]int `json:"errors_by_type"`
	StartTime      time.Time      `json:"start_time"`
	EndTime        time.Time      `json:"end_time"`
}

var (
	currentTest *TestSummary
	testMutex   sync.RWMutex
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", statusPageHandler).Methods("GET")
	r.HandleFunc("/status", statusHandler).Methods("GET")
	r.HandleFunc("/test/{concurrency}", testHandler).Methods("GET")

	log.Println("Load service starting on port 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}

func statusPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write(
		[]byte(`
<!DOCTYPE html>
<html>
<head>
    <title>Load Testing Service (RPS)</title>
</head>
<body>
    <h1>Load Testing Service (RPS)</h1>
    <p>Status: Running</p>
    <p>Available endpoints:</p>
    <ul>
        <li><a href="/status">GET /status</a> - Current test status</li>
        <li>GET /test/{rps} - Run load test with specified requests per second</li>
        <li>GET /test/{rps}?duration={seconds} - Run load test with custom duration (default: 30s)</li>
    </ul>
    <p>Example tests:</p>
    <ul>
        <li><a href="/test/1">Low load (1 RPS)</a></li>
        <li><a href="/test/5">Medium load (5 RPS)</a></li>
        <li><a href="/test/10">High load (10 RPS)</a></li>
        <li><a href="/test/20">Overload (20 RPS)</a></li>
        <li><a href="/test/50">Extreme load (50 RPS)</a></li>

    </ul>
    <p>Custom duration examples:</p>
    <ul>
        <li><a href="/test/10?duration=20">10 RPS for 20 seconds</a></li>
        <li><a href="/test/20?duration=20">20 RPS for 20 seconds</a></li>
        <li><a href="/test/30?duration=20">30 RPS for 20 seconds</a></li>
        <li><a href="/test/50?duration=20">50 RPS for 20 seconds</a></li>
        <li><a href="/test/100?duration=20">100 RPS for 20 seconds</a></li>
        <li><a href="/test/200?duration=20">200 RPS for 20 seconds</a></li>
    </ul>
</body>
</html>
	`),
	)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	testMutex.RLock()
	defer testMutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")

	if currentTest == nil {
		json.NewEncoder(w).Encode(
			map[string]string{
				"status": "No tests run yet",
			},
		)
		return
	}

	json.NewEncoder(w).Encode(currentTest)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	rpsParam := mux.Vars(r)["concurrency"] // keeping same URL param name for compatibility
	rps, err := strconv.Atoi(rpsParam)
	if err != nil {
		http.Error(w, "Invalid RPS value", 400)
		return
	}

	if rps <= 0 || rps > 1000 {
		http.Error(w, "RPS must be between 1 and 1000", 400)
		return
	}

	// Default test duration is 30 seconds
	testDuration := 30 * time.Second
	if durationParam := r.URL.Query().Get("duration"); durationParam != "" {
		if d, err := strconv.Atoi(durationParam); err == nil && d > 0 && d <= 300 {
			testDuration = time.Duration(d) * time.Second
		}
	}

	results := make(chan TestResult, rps*int(testDuration.Seconds())+100) // buffer for all possible requests
	var wg sync.WaitGroup

	startTime := time.Now()
	ticker := time.NewTicker(time.Second / time.Duration(rps))
	defer ticker.Stop()

	testCtx, cancel := context.WithTimeout(context.Background(), testDuration)
	defer cancel()

	requestNum := 0

	// Launch requests at specified RPS
	for {
		select {
		case <-testCtx.Done():
			// Test duration completed
			goto waitForCompletion
		case <-ticker.C:
			wg.Add(1)
			currentReqNum := requestNum
			requestNum++

			go func(reqNum int) {
				defer wg.Done()

				result := TestResult{
					RequestNum: reqNum,
					StartTime:  time.Now(),
				}

				client := &http.Client{}

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				req, err := http.NewRequestWithContext(
					ctx,
					"POST",
					"http://tx-service:8081/process",
					strings.NewReader("{}"),
				)
				if err != nil {
					result.Error = err.Error()
					result.Success = false
					result.EndTime = time.Now()
					result.Duration = result.EndTime.Sub(result.StartTime)
					results <- result
					return
				}
				req.Header.Set("Content-Type", "application/json")

				resp, err := client.Do(req)

				result.EndTime = time.Now()
				result.Duration = result.EndTime.Sub(result.StartTime)

				if err != nil {
					fmt.Println("Error processing request:", err)
					result.Error = err.Error()
					result.Success = false
				} else {
					result.StatusCode = resp.StatusCode
					result.Success = resp.StatusCode == 200
					resp.Body.Close()
				}

				results <- result
			}(currentReqNum)
		}
	}

waitForCompletion:
	// Wait for all requests to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Aggregate results
	summary := aggregateResults(results, startTime, requestNum)

	// Store current test results
	testMutex.Lock()
	currentTest = &summary
	testMutex.Unlock()

	log.Printf(
		"Load service: test completed - %d/%d succeeded (%.1f RPS) - Avg latency: %.2fs - Error breakdown: %v",
		summary.SuccessCount, summary.TotalRequests, float64(summary.TotalRequests)/testDuration.Seconds(),
		summary.AverageLatency, summary.ErrorsByType,
	)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

func aggregateResults(results <-chan TestResult, startTime time.Time, totalRequests int) TestSummary {
	var allResults []TestResult
	errorCounts := make(map[string]int)
	successCount := 0

	for result := range results {
		allResults = append(allResults, result)

		if result.Success {
			successCount++
		} else {
			errorType := categorizeError(result.Error, result.StatusCode)
			errorCounts[errorType]++
		}
	}

	// Calculate latency statistics
	sort.Slice(
		allResults, func(i, j int) bool {
			return allResults[i].Duration < allResults[j].Duration
		},
	)

	var totalLatency time.Duration
	for _, result := range allResults {
		totalLatency += result.Duration
	}

	endTime := time.Now()

	summary := TestSummary{
		TotalRequests:  totalRequests,
		SuccessCount:   successCount,
		FailureCount:   len(allResults) - successCount,
		SuccessRate:    float64(successCount) / float64(len(allResults)) * 100,
		TotalDuration:  endTime.Sub(startTime).Seconds(),
		AverageLatency: (totalLatency / time.Duration(len(allResults))).Seconds(),
		ErrorsByType:   errorCounts,
		StartTime:      startTime,
		EndTime:        endTime,
	}

	if len(allResults) > 0 {
		summary.MinLatency = allResults[0].Duration.Seconds()
		summary.MaxLatency = allResults[len(allResults)-1].Duration.Seconds()

		// Calculate P95
		p95Index := int(float64(len(allResults)) * 0.95)
		if p95Index < len(allResults) {
			summary.P95Latency = allResults[p95Index].Duration.Seconds()
		}
	}

	return summary
}

func categorizeError(errorMsg string, statusCode int) string {
	if statusCode == 503 {
		return "connection_pool_exhausted"
	}
	if statusCode == 500 {
		return "internal_server_error"
	}
	if strings.Contains(errorMsg, "connection refused") {
		return "connection_refused"
	}
	if strings.Contains(errorMsg, "timeout") {
		return "timeout"
	}
	return "other"
}
