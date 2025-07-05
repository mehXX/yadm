# PostgreSQL Connection Exhaustion Demonstration

## Overall Architecture Plan

**Project Structure:**
```
01_postgres_tx_connections/
├── docker-compose.yml          # PostgreSQL + services orchestration
├── slow-service/              # Service 1: Slow HTTP endpoint
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── tx-service/                # Service 2: Transaction + external call
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── load-service/              # Service 3: Load generator
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
└── postgres/                  # PostgreSQL configuration
    └── init.sql
```

**Architecture Flow:**
1. **PostgreSQL**: Default config (100 connections max)
2. **Slow Service**: HTTP server with 5-second artificial delay
3. **TX Service**: Opens DB transaction → calls slow service → commits
4. **Load Service**: Sends concurrent requests to TX service
5. **Observability**: Structured logging to demonstrate connection exhaustion

**Demonstration Steps:**
1. Start with low concurrency (5 requests) - everything works
2. Increase to medium concurrency (50 requests) - slower responses
3. Increase to high concurrency (150+ requests) - connection failures

**Expected Demonstration:**
- **Low concurrency (5 requests)**: All succeed, ~5 second response time
- **Medium concurrency (50 requests)**: Some queueing, longer response times
- **High concurrency (150 requests)**: Connection pool exhaustion, "too many connections" errors

**Key Failure Points to Observe:**
- PostgreSQL connection limit reached
- Transactions holding connections during slow external calls
- Request failures with connection errors
- Performance degradation under load

**Usage Instructions:**
1. `docker-compose up` to start all services
2. Test with `curl http://localhost:8082/test/5` (low load)
3. Increase load: `curl http://localhost:8082/test/150` (high load)
4. Monitor logs for connection errors