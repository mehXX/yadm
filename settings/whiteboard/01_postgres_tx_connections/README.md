# PostgreSQL Transaction Connections Demo

A demonstration of PostgreSQL transaction handling and connection management using Go microservices.

## Services

- **postgres**: PostgreSQL database with connection logging
- **slow-service**: Simulates slow operations (port 8080)
- **tx-service**: Handles transactions and database operations (port 8081)
- **load-service**: Generates load and tests the system (port 8082)

## Quick Start

1. **Start all services:**
   ```bash
   docker-compose up -d
   ```

2. **Check service status:**
   ```bash
   docker-compose ps
   ```

3. **View logs:**
   ```bash
   # All services
   docker-compose logs -f
   
   # Specific service
   docker-compose logs -f postgres
   docker-compose logs -f tx-service
   ```

## Testing

The services expose the following endpoints:

- **Load Service**: http://localhost:8082
- **TX Service**: http://localhost:8081  
- **Slow Service**: http://localhost:8080

## Database Access

Connect to PostgreSQL directly:
```bash
docker exec -it postgres-tx-demo psql -U testuser -d testdb
```

## Cleanup

```bash
docker-compose down
docker-compose down -v  # Remove volumes too
```

## Connection Monitoring

PostgreSQL is configured with extensive connection logging. Check the postgres logs to monitor connection behavior:
```bash
docker-compose logs -f postgres
```