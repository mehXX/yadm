# PostgreSQL Transaction Connection Exhaustion Demo

## Project Purpose
Demonstrates the PostgreSQL connection exhaustion problem when making external HTTP calls within database transactions. Shows how long-running transactions can exhaust connection pools and cause cascading failures.

## Architecture Overview

### System Components
```
01_postgres_tx_connections/
├── docker-compose.yml          # Container orchestration
├── postgres/
│   └── init.sql               # Database initialization
├── slow-service/              # HTTP endpoint with artificial delay
│   ├── main.go               # 3-second delay simulation
│   ├── go.mod
│   └── Dockerfile
├── tx-service/                # Transaction service (main bottleneck)
│   ├── main.go               # DB transaction + external HTTP call
│   ├── go.mod
│   └── Dockerfile
└── load-service/              # RPS-based load generator
    ├── main.go               # Configurable requests per second
    ├── go.mod
    └── Dockerfile
```

### Service Details

#### 1. PostgreSQL Database
- **Image**: postgres:15
- **Connection Limit**: Default (100 connections)
- **Database**: testdb
- **Tables**: 
  - `demo_table` (from init.sql)
  - `requests` (created by tx-service)

#### 2. Slow Service (Port 8080)
- **Purpose**: Simulates slow external API
- **Delay**: 3 seconds per request
- **Metrics**: Tracks total/active requests
- **Endpoints**:
  - `GET /slow` - Returns JSON after 3s delay

#### 3. TX Service (Port 8081)
- **Purpose**: Core demonstration service
- **Database**: pgx connection pool (max 100 connections)
- **Flow**: 
  1. Begin transaction
  2. Call slow-service (3s delay)
  3. Insert record to `requests` table
  4. Commit transaction
- **Endpoints**:
  - `POST /process` - Main transaction endpoint
  - `GET /health` - Connection pool status

#### 4. Load Service (Port 8082)
- **Purpose**: RPS-based load testing
- **Method**: Controlled requests per second
- **Metrics**: Success rate, latency percentiles, error categorization
- **Endpoints**:
  - `GET /` - Web interface
  - `GET /status` - Current test results
  - `GET /test/{rps}` - Run test at specified RPS
  - `GET /test/{rps}?duration={seconds}` - Custom duration


## Key Observations
1. **Connection Exhaustion**: Occurs around 35 RPS
2. **Error Patterns**: 503 errors when pool is exhausted
3. **Latency Impact**: Response times increase before failures
4. **Recovery**: System recovers when load decreases