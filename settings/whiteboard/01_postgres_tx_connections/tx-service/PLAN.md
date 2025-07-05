# TX Service Plan

## Purpose
Demonstrate connection exhaustion by holding DB connections during slow external calls

## Technical Details
- **Port**: 8081
- **Endpoint**: `POST /process`
- **Database**: PostgreSQL with PGX driver
- **Critical Flow**: Begin TX → Call slow service → Commit TX

## Implementation Strategy

### Main Handler
```go
func processHandler(w http.ResponseWriter, r *http.Request) {
    requestID := generateRequestID()
    log.Printf("TX service: starting request %s", requestID)
    
    // 1. Begin transaction (holds connection)
    tx, err := db.Begin(context.Background())
    if err != nil {
        log.Printf("TX service: failed to begin transaction: %v", err)
        http.Error(w, "Connection pool exhausted", 503)
        return
    }
    defer tx.Rollback(context.Background())
    
    // 2. Call slow external service (connection still held)
    slowResponse, err := callSlowService(requestID)
    if err != nil {
        log.Printf("TX service: slow service call failed: %v", err)
        http.Error(w, "External service error", 500)
        return
    }
    
    // 3. Insert record in transaction
    _, err = tx.Exec(context.Background(), 
        "INSERT INTO requests (request_id, external_call_duration) VALUES ($1, $2)",
        requestID, 5000)
    if err != nil {
        log.Printf("TX service: insert failed: %v", err)
        http.Error(w, "Database insert failed", 500)
        return
    }
    
    // 4. Commit transaction
    if err := tx.Commit(context.Background()); err != nil {
        log.Printf("TX service: commit failed: %v", err)
        http.Error(w, "Transaction commit failed", 500)
        return
    }
    
    log.Printf("TX service: completed request %s", requestID)
    json.NewEncoder(w).Encode(slowResponse)
}
```

### External Service Call
```go
func callSlowService(requestID string) (map[string]interface{}, error) {
    log.Printf("TX service: calling slow service for request %s", requestID)
    
    resp, err := http.Get("http://slow-service:8080/slow")
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
```

### Database Connection Setup
```go
func setupDatabase() (*pgxpool.Pool, error) {
    config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
    if err != nil {
        return nil, fmt.Errorf("failed to parse database URL: %w", err)
    }
    
    // Configure connection pool
    config.MaxConns = 30        // Limit connections from this service
    config.MinConns = 5
    config.MaxConnLifetime = time.Hour
    config.MaxConnIdleTime = time.Minute * 30
    
    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        return nil, fmt.Errorf("failed to create connection pool: %w", err)
    }
    
    return pool, nil
}
```

## Database Schema Requirements
```sql
CREATE TABLE IF NOT EXISTS requests (
    id SERIAL PRIMARY KEY,
    request_id VARCHAR(255) NOT NULL,
    processed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    external_call_duration INTEGER -- milliseconds
);
```

## Logging Requirements
- Transaction begin/commit status
- Connection pool statistics
- External service call duration
- Failed connections count
- Request processing time

## Expected Behavior
- Each request holds a database connection for ~5 seconds
- Under high load, connection pool gets exhausted
- New requests fail with "too many connections" error
- Successful requests insert record and return slow service response

## Files to Create
1. `main.go` - HTTP server with transaction logic
2. `go.mod` - Go module with PGX dependency
3. `Dockerfile` - Container configuration

## Environment Variables
- `DATABASE_URL`: PostgreSQL connection string
- `SLOW_SERVICE_URL`: URL of slow service (default: http://slow-service:8080)

## Testing
- Single request: `curl -X POST http://localhost:8081/process`
- Load test: Use load service to send concurrent requests
- Monitor PostgreSQL connection count during load