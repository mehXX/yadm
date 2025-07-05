# Slow HTTP Service Plan

## Purpose
Simulate a slow external API that takes 5 seconds to respond

## Technical Details
- **Port**: 8080
- **Endpoint**: `GET /slow`
- **Response Time**: Fixed 5 seconds using `time.Sleep(5 * time.Second)`
- **Response**: Simple JSON with timestamp and request ID

## Implementation Strategy

### HTTP Handler
```go
func slowHandler(w http.ResponseWriter, r *http.Request) {
    requestID := generateRequestID()
    log.Printf("Slow service: received request %s", requestID)
    
    time.Sleep(5 * time.Second)
    
    response := map[string]interface{}{
        "request_id": requestID,
        "timestamp": time.Now(),
        "message": "Response after 5 seconds",
    }
    
    json.NewEncoder(w).Encode(response)
    log.Printf("Slow service: completed request %s", requestID)
}
```

### Request ID Generation
```go
func generateRequestID() string {
    return fmt.Sprintf("req_%d_%s", 
        time.Now().UnixNano(), 
        randomString(6))
}
```

## Logging Requirements
- Request received with request ID
- Request completed with request ID
- Metrics tracking:
  - Total requests received
  - Requests per second
  - Active requests count

## Expected Behavior
- Every request takes exactly 5 seconds
- Concurrent requests are handled independently
- Each request gets unique ID for tracking
- Logs allow correlation of request start/end

## Files to Create
1. `main.go` - HTTP server with slow endpoint
2. `go.mod` - Go module definition
3. `Dockerfile` - Container configuration

## Testing
- Single request: `curl http://localhost:8080/slow`
- Multiple requests: Use load service to test concurrency
- Verify 5-second response time consistently