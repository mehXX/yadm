package main

import (
	"encoding/json"
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
	TotalRequests    int               `json:"total_requests"`
	SuccessCount     int               `json:"success_count"`
	FailureCount     int               `json:"failure_count"`
	SuccessRate      float64           `json:"success_rate"`
	TotalDuration    time.Duration     `json:"total_duration_ms"`
	AverageLatency   time.Duration     `json:"average_latency_ms"`
	MinLatency       time.Duration     `json:"min_latency_ms"`
	MaxLatency       time.Duration     `json:"max_latency_ms"`
	P95Latency       time.Duration     `json:"p95_latency_ms"`
	ErrorsByType     map[string]int    `json:"errors_by_type"`
	StartTime        time.Time         `json:"start_time"`
	EndTime          time.Time         `json:"end_time"`
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
	w.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
    <title>Load Testing Service</title>
</head>
<body>
    <h1>Load Testing Service</h1>
    <p>Status: Running</p>
    <p>Available endpoints:</p>
    <ul>
        <li><a href="/status">GET /status</a> - Current test status</li>
        <li>GET /test/{concurrency} - Run load test with specified concurrency</li>
    </ul>
    <p>Example tests:</p>
    <ul>
        <li><a href="/test/5">Low load (5 concurrent requests)</a></li>
        <li><a href="/test/25">Medium load (25 concurrent requests)</a></li>
        <li><a href="/test/50">High load (50 concurrent requests)</a></li>
        <li><a href="/test/100">Overload (100 concurrent requests)</a></li>
        <li><a href="/test/150">Extreme load (150 concurrent requests)</a></li>
    </ul>
</body>
</html>
	`))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	testMutex.RLock()
	defer testMutex.RUnlock()
	
	w.Header().Set("Content-Type", "application/json")
	
	if currentTest == nil {
		json.NewEncoder(w).Encode(map[string]string{
			"status": "No tests run yet",
		})
		return
	}
	
	json.NewEncoder(w).Encode(currentTest)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	concurrency := mux.Vars(r)["concurrency"]
	level, err := strconv.Atoi(concurrency)
	if err != nil {
		http.Error(w, "Invalid concurrency level", 400)
		return
	}
	
	if level <= 0 || level > 1000 {
		http.Error(w, "Concurrency level must be between 1 and 1000", 400)
		return
	}
	
	log.Printf("Load service: starting test with %d concurrent requests", level)
	
	results := make(chan TestResult, level)
	var wg sync.WaitGroup
	
	startTime := time.Now()
	
	// Launch concurrent requests
	for i := 0; i < level; i++ {
		wg.Add(1)
		go func(requestNum int) {
			defer wg.Done()
			
			result := TestResult{
				RequestNum: requestNum,
				StartTime:  time.Now(),
			}
			
			resp, err := http.Post("http://tx-service:8081/process", 
				"application/json", strings.NewReader("{}"))
			
			result.EndTime = time.Now()
			result.Duration = result.EndTime.Sub(result.StartTime)
			
			if err != nil {
				result.Error = err.Error()
				result.Success = false
			} else {
				result.StatusCode = resp.StatusCode
				result.Success = resp.StatusCode == 200
				resp.Body.Close()
			}
			
			results <- result
		}(i)
	}
	
	// Wait and collect results
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Aggregate results
	summary := aggregateResults(results, startTime, level)
	
	// Store current test results
	testMutex.Lock()
	currentTest = &summary
	testMutex.Unlock()
	
	log.Printf("Load service: test completed - %d/%d succeeded", 
		summary.SuccessCount, summary.TotalRequests)
	
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
	sort.Slice(allResults, func(i, j int) bool {
		return allResults[i].Duration < allResults[j].Duration
	})
	
	var totalLatency time.Duration
	for _, result := range allResults {
		totalLatency += result.Duration
	}
	
	endTime := time.Now()
	
	summary := TestSummary{
		TotalRequests:    totalRequests,
		SuccessCount:     successCount,
		FailureCount:     len(allResults) - successCount,
		SuccessRate:      float64(successCount) / float64(len(allResults)) * 100,
		TotalDuration:    endTime.Sub(startTime),
		AverageLatency:   totalLatency / time.Duration(len(allResults)),
		ErrorsByType:     errorCounts,
		StartTime:        startTime,
		EndTime:          endTime,
	}
	
	if len(allResults) > 0 {
		summary.MinLatency = allResults[0].Duration
		summary.MaxLatency = allResults[len(allResults)-1].Duration
		
		// Calculate P95
		p95Index := int(float64(len(allResults)) * 0.95)
		if p95Index < len(allResults) {
			summary.P95Latency = allResults[p95Index].Duration
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