# Docker Setup Guide for Zlagoda

## Quick Setup

### 1. Clone and Setup Environment
```bash
git clone <repository-url>
cd zlagoda
cp .env.sample .env
```

### 2. Configure Environment Variables
Edit `.env` file with your settings:

```env
# Database Configuration
DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=zlagoda_user
DB_PASSWORD=your_secure_password
DB_NAME=zlagoda_db

# Server Configuration
PORT=8080

# Business Configuration
VAT_RATE=0.2
```

### 3. Start Services
```bash
sudo docker compose up --build
```

## Troubleshooting Common Issues

### Issue 1: Migration Connection Refused
**Error**: `dial tcp [::1]:5432: connect: connection refused`

**Solution**: This happens when the migrator tries to connect before PostgreSQL is ready.

1. **Check if containers are running**:
   ```bash
   docker ps
   ```

2. **Check PostgreSQL logs**:
   ```bash
   docker logs postgres-zlagoda
   ```

3. **Restart with proper wait**:
   ```bash
   docker compose down
   docker compose up --build
   ```

### Issue 2: Database Connection Issues
**Error**: Various database connection errors

**Solutions**:

1. **Verify environment variables**:
   ```bash
   # Check .env file exists and has correct values
   cat .env
   ```

2. **Check database container health**:
   ```bash
   docker inspect postgres-zlagoda --format='{{.State.Health.Status}}'
   ```

3. **Manual database connection test**:
   ```bash
   docker exec -it postgres-zlagoda psql -U zlagoda_user -d zlagoda_db
   ```

### Issue 3: Port Already in Use
**Error**: `port is already allocated`

**Solutions**:

1. **Check what's using the port**:
   ```bash
   sudo lsof -i :8080
   sudo lsof -i :5432
   ```

2. **Kill processes or change ports in `.env`**:
   ```bash
   # Kill process using port
   sudo kill -9 <PID>
   
   # Or change PORT in .env file
   PORT=8081
   ```

### Issue 4: Build Failures
**Error**: Build context or dependency issues

**Solutions**:

1. **Clean Docker cache**:
   ```bash
   docker system prune -a
   docker volume prune
   ```

2. **Rebuild without cache**:
   ```bash
   docker compose build --no-cache
   docker compose up
   ```

## Manual Setup Steps

### 1. Run Each Service Separately

#### Start PostgreSQL:
```bash
docker compose up postgres-zlagoda
```

#### Run Migrations (in new terminal):
```bash
docker compose up migrator-zlagoda
```

#### Start API (after migrations complete):
```bash
docker compose up api-zlagoda
```

### 2. Development Mode (API outside Docker)

#### Start only database:
```bash
docker compose up postgres-zlagoda
```

#### Update .env for local development:
```env
DB_HOST=localhost
DB_PORT=5432
```

#### Run migrations locally:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
./migrate.sh
```

#### Run API locally:
```bash
go run cmd/api/main.go
```

## Verifying Setup

### 1. Check Service Health
```bash
# Check all containers
docker ps

# Check API health
curl http://localhost:8080/api/categories

# Check database connection
docker exec -it postgres-zlagoda psql -U zlagoda_user -d zlagoda_db -c "\dt"
```

### 2. Test API Endpoints
```bash
# Create a category
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{"name": "Test Category"}'

# List categories
curl http://localhost:8080/api/categories

# Create a product
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Test Product", "characteristics": "Test characteristics", "category_id": 1}'
```

## Environment Variables Reference

### Required Variables
- `DB_DRIVER`: Database driver (always `postgres`)
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `PORT`: API server port

### Docker-Specific Overrides
When running in Docker, these are automatically set:
- `DB_HOST`: Set to `postgres-zlagoda` (service name)
- `DB_PORT`: Set to `5432` (internal container port)

### Optional Variables
- `VAT_RATE`: Tax rate (default: 0.2 = 20%)

## Common Commands

### Container Management
```bash
# View logs
docker compose logs api-zlagoda
docker compose logs postgres-zlagoda
docker compose logs migrator-zlagoda

# Restart specific service
docker compose restart api-zlagoda

# Stop all services
docker compose down

# Remove volumes (clears database)
docker compose down -v
```

### Database Management
```bash
# Connect to database
docker exec -it postgres-zlagoda psql -U zlagoda_user -d zlagoda_db

# View tables
\dt

# View table structure
\d category

# Exit psql
\q
```

### Development Commands
```bash
# Build and test locally
go build cmd/api/main.go
go test ./...

# Run with local database
DB_HOST=localhost go run cmd/api/main.go
```

## Production Considerations

### Security
1. **Change default passwords** in production
2. **Use Docker secrets** for sensitive data
3. **Enable SSL/TLS** for database connections
4. **Restrict network access** to database

### Performance
1. **Configure connection pooling** via environment variables:
   ```env
   DB_MAX_OPEN_CONNS=25
   DB_MAX_IDLE_CONNS=5
   DB_CONN_MAX_LIFETIME=5m
   ```

2. **Use external database** for production:
   ```env
   DB_HOST=your-production-db-host
   ```

### Monitoring
1. **Add health check endpoints**
2. **Configure logging** with proper levels
3. **Set up metrics collection**
4. **Monitor container resource usage**

## Need Help?

1. **Check logs first**: `docker compose logs`
2. **Verify environment**: `cat .env`
3. **Test database connection**: `docker exec -it postgres-zlagoda psql -U zlagoda_user -d zlagoda_db`
4. **Clean restart**: `docker compose down && docker compose up --build`

For more detailed API documentation, see [README.md](README.md).