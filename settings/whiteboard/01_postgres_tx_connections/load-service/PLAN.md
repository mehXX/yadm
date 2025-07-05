# Load Testing Service Plan

## Purpose
Generate concurrent load to demonstrate connection exhaustion

## Technical Details
- **Port**: 8082
- **Endpoints**: 
  - `GET /test/{concurrency}` - Run test with specified concurrency level
  - `GET /status` - Current test status and metrics
  - `GET /` - Simple status page
- **Load Patterns**: 5, 25, 50, 100, 150 concurrent requests

## Implementation Strategy

### Main Test Handler
```go
func testHandler(w http.ResponseWriter, r *http.Request) {
    concurrency := mux.Vars(r)["concurrency"]
    level, err := strconv.Atoi(concurrency)
    if err != nil {
        http.Error(w, "Invalid concurrency level", 400)
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
    
    log.Printf("Load service: test completed - %d/%d succeeded", 
        summary.SuccessCount, summary.TotalRequests)
    
    json.NewEncoder(w).Encode(summary)
}
```

### Test Result Structures
```go
type TestResult struct {
    RequestNum  int           `json:"request_num"`
    StartTime   time.Time     `json:"start_time"`
    EndTime     time.Time     `json:"end_time"`
    Duration    time.Duration `json:"duration_ms"`
    StatusCode  int           `json:"status_code"`
    Success     bool          `json:"success"`
    Error       string        `json:"error,omitempty"`
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
```

### Results Aggregation
```go
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
```

### Error Categorization
```go
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
```

## Test Scenarios
1. **Low Load**: 5 concurrent requests (should succeed)
2. **Medium Load**: 25 concurrent requests (may start showing delays)
3. **High Load**: 50 concurrent requests (some failures likely)
4. **Overload**: 100 concurrent requests (connection issues likely)
5. **Extreme Load**: 150 concurrent requests (definite connection exhaustion)

## Metrics Collection
- Success/failure rates
- Response times (min, max, avg, p95)
- Connection errors count by type
- Total test duration
- Requests per second

## Expected Behavior
- Low concurrency: All requests succeed
- High concurrency: Connection pool exhaustion errors
- Clear correlation between concurrency and failure rate
- Detailed error categorization for analysis

## Files to Create
1. `main.go` - HTTP server with load testing logic
2. `go.mod` - Go module definition
3. `Dockerfile` - Container configuration

## Usage Examples
```bash
# Test with different concurrency levels
curl http://localhost:8082/test/5     # Low load
curl http://localhost:8082/test/25    # Medium load
curl http://localhost:8082/test/50    # High load
curl http://localhost:8082/test/100   # Overload
curl http://localhost:8082/test/150   # Extreme load

# Check current status
curl http://localhost:8082/status
```

## Logging Requirements
- Test start/completion with parameters
- Real-time progress updates
- Error summaries by type
- Performance metrics