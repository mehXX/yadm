# Docker Compose Configuration Plan

## Purpose
Orchestrate all services and PostgreSQL to demonstrate connection exhaustion

## Service Architecture
- **postgres**: PostgreSQL with connection limits and logging
- **slow-service**: HTTP service with artificial delay
- **tx-service**: Transaction service that calls slow-service
- **load-service**: Load generator for testing

## Complete Docker Compose Configuration

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15
    container_name: postgres-tx-demo
    environment:
      POSTGRES_DB: testdb
      POSTGRES_USER: testuser
      POSTGRES_PASSWORD: testpass
    ports:
      - "5432:5432"
    command: |
      postgres 
      -c max_connections=100
      -c shared_preload_libraries=pg_stat_statements
      -c log_connections=on
      -c log_disconnections=on
      -c log_statement=all
      -c log_min_duration_statement=1000
      -c log_line_prefix='%t [%p]: [%l-1] user=%u,db=%d,app=%a,client=%h '
    volumes:
      - ./postgres/init.sql:/docker-entrypoint-initdb.d/init.sql
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U testuser -d testdb"]
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - demo-network

  slow-service:
    build: 
      context: ./slow-service
      dockerfile: Dockerfile
    container_name: slow-service-demo
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      - LOG_LEVEL=info
    depends_on:
      postgres:
        condition: service_healthy
    networks:
      - demo-network
    restart: unless-stopped

  tx-service:
    build: 
      context: ./tx-service
      dockerfile: Dockerfile
    container_name: tx-service-demo
    ports:
      - "8081:8081"
    environment:
      - PORT=8081
      - DATABASE_URL=postgres://testuser:testpass@postgres:5432/testdb
      - SLOW_SERVICE_URL=http://slow-service:8080
      - LOG_LEVEL=info
    depends_on:
      postgres:
        condition: service_healthy
      slow-service:
        condition: service_started
    networks:
      - demo-network
    restart: unless-stopped

  load-service:
    build: 
      context: ./load-service
      dockerfile: Dockerfile
    container_name: load-service-demo
    ports:
      - "8082:8082"
    environment:
      - PORT=8082
      - TX_SERVICE_URL=http://tx-service:8081
      - LOG_LEVEL=info
    depends_on:
      tx-service:
        condition: service_started
    networks:
      - demo-network
    restart: unless-stopped

volumes:
  postgres_data:
    driver: local

networks:
  demo-network:
    driver: bridge
```

## Service Dependencies
1. **postgres**: Base service, must be healthy before others start
2. **slow-service**: Depends on postgres being healthy
3. **tx-service**: Depends on both postgres and slow-service
4. **load-service**: Depends on tx-service

## Port Mapping
- **5432**: PostgreSQL database
- **8080**: Slow HTTP service
- **8081**: Transaction service
- **8082**: Load testing service

## Environment Variables

### PostgreSQL Service
- `POSTGRES_DB`: Database name (testdb)
- `POSTGRES_USER`: Database user (testuser)
- `POSTGRES_PASSWORD`: Database password (testpass)

### Slow Service
- `PORT`: HTTP server port (8080)
- `LOG_LEVEL`: Logging verbosity (info)

### TX Service
- `PORT`: HTTP server port (8081)
- `DATABASE_URL`: PostgreSQL connection string
- `SLOW_SERVICE_URL`: URL of slow service
- `LOG_LEVEL`: Logging verbosity (info)

### Load Service
- `PORT`: HTTP server port (8082)
- `TX_SERVICE_URL`: URL of transaction service
- `LOG_LEVEL`: Logging verbosity (info)

## Health Checks
- **postgres**: `pg_isready` command to ensure database is accepting connections
- **Other services**: Implicit dependency on postgres health

## Networking
- **Custom network**: `demo-network` for inter-service communication
- **Service discovery**: Services can reach each other by service name

## Volumes
- **postgres_data**: Persistent storage for PostgreSQL data
- **init.sql**: Database initialization script

## Usage Commands

### Start Services
```bash
# Start all services
docker-compose up -d

# Start with logs
docker-compose up

# Start specific service
docker-compose up postgres
```

### Monitor Services
```bash
# View logs
docker-compose logs -f
docker-compose logs postgres
docker-compose logs tx-service

# Check service status
docker-compose ps

# View resource usage
docker-compose top
```

### Testing Commands
```bash
# Test slow service
curl http://localhost:8080/slow

# Test transaction service
curl -X POST http://localhost:8081/process

# Run load tests
curl http://localhost:8082/test/5    # Low load
curl http://localhost:8082/test/50   # Medium load
curl http://localhost:8082/test/100  # High load
curl http://localhost:8082/test/150  # Overload
```

### Database Monitoring
```bash
# Connect to database
docker-compose exec postgres psql -U testuser -d testdb

# Monitor connections
docker-compose exec postgres psql -U testuser -d testdb -c "SELECT * FROM connection_stats;"

# Check recent requests
docker-compose exec postgres psql -U testuser -d testdb -c "SELECT * FROM request_stats;"
```

### Cleanup
```bash
# Stop services
docker-compose down

# Stop and remove volumes
docker-compose down -v

# Remove everything including images
docker-compose down -v --rmi all
```

## Expected Demonstration Flow
1. **Start**: `docker-compose up -d`
2. **Low Load**: `curl http://localhost:8082/test/5` (should succeed)
3. **Monitor**: Watch connection count in postgres logs
4. **High Load**: `curl http://localhost:8082/test/150` (should fail)
5. **Observe**: Connection exhaustion errors in logs
6. **Cleanup**: `docker-compose down`

## Logging Strategy
- **Structured logging**: JSON format for easy parsing
- **Service identification**: Each service prefixes logs with service name
- **Connection tracking**: PostgreSQL logs all connections/disconnections
- **Request correlation**: Request IDs for tracing across services

## Troubleshooting
- **Connection refused**: Check if postgres is healthy
- **Service not found**: Verify network configuration
- **Port conflicts**: Ensure ports 5432, 8080, 8081, 8082 are available
- **Permission denied**: Check file permissions on init.sql