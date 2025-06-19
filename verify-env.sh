#!/bin/bash

echo "=== Zlagoda Environment Verification ==="
echo

# Check if .env file exists
if [ ! -f .env ]; then
    echo "❌ ERROR: .env file not found!"
    echo "   Please copy .env.sample to .env and configure it:"
    echo "   cp .env.sample .env"
    exit 1
fi

echo "✅ .env file found"

# Required environment variables
REQUIRED_VARS=(
    "DB_DRIVER"
    "DB_USER"
    "DB_PASSWORD"
    "DB_NAME"
    "PORT"
)

OPTIONAL_VARS=(
    "VAT_RATE"
    "DB_HOST"
    "DB_PORT"
)

# Load environment variables
source .env

echo
echo "=== Required Variables ==="

missing_vars=0
for var in "${REQUIRED_VARS[@]}"; do
    if [ -z "${!var}" ]; then
        echo "❌ $var is not set or empty"
        missing_vars=$((missing_vars + 1))
    else
        # Mask password for display
        if [ "$var" = "DB_PASSWORD" ]; then
            echo "✅ $var = [HIDDEN]"
        else
            echo "✅ $var = ${!var}"
        fi
    fi
done

echo
echo "=== Optional Variables ==="

for var in "${OPTIONAL_VARS[@]}"; do
    if [ -z "${!var}" ]; then
        case $var in
            "VAT_RATE")
                echo "⚠️  $var not set (will use default: 0.2)"
                ;;
            "DB_HOST")
                echo "⚠️  $var not set (will use localhost for local dev, postgres-zlagoda in Docker)"
                ;;
            "DB_PORT")
                echo "⚠️  $var not set (will use 5432)"
                ;;
        esac
    else
        echo "✅ $var = ${!var}"
    fi
done

echo
echo "=== Configuration Validation ==="

# Validate DB_DRIVER
if [ "$DB_DRIVER" != "postgres" ]; then
    echo "❌ DB_DRIVER should be 'postgres', got: $DB_DRIVER"
    missing_vars=$((missing_vars + 1))
else
    echo "✅ DB_DRIVER is correct"
fi

# Validate PORT
if [ -n "$PORT" ]; then
    if ! [[ "$PORT" =~ ^[0-9]+$ ]] || [ "$PORT" -lt 1 ] || [ "$PORT" -gt 65535 ]; then
        echo "❌ PORT should be a number between 1-65535, got: $PORT"
        missing_vars=$((missing_vars + 1))
    else
        echo "✅ PORT is valid"
    fi
fi

# Validate VAT_RATE if set
if [ -n "$VAT_RATE" ]; then
    if ! [[ "$VAT_RATE" =~ ^[0-9]*\.?[0-9]+$ ]] || (( $(echo "$VAT_RATE < 0" | bc -l) )) || (( $(echo "$VAT_RATE > 1" | bc -l) )); then
        echo "❌ VAT_RATE should be a decimal between 0-1, got: $VAT_RATE"
        missing_vars=$((missing_vars + 1))
    else
        echo "✅ VAT_RATE is valid"
    fi
fi

echo
echo "=== Docker Environment Check ==="

# Check if Docker is available
if command -v docker &> /dev/null; then
    echo "✅ Docker is installed"

    # Check if Docker Compose is available
    if docker compose version &> /dev/null; then
        echo "✅ Docker Compose is available"
    else
        echo "❌ Docker Compose is not available"
        echo "   Please install Docker Compose"
        missing_vars=$((missing_vars + 1))
    fi
else
    echo "❌ Docker is not installed"
    echo "   Please install Docker to use containerized setup"
fi

echo
echo "=== Port Availability Check ==="

# Check if ports are available
if [ -n "$PORT" ]; then
    if lsof -Pi :$PORT -sTCP:LISTEN -t >/dev/null ; then
        echo "⚠️  Port $PORT is already in use"
        echo "   You may need to stop the service using this port or change PORT in .env"
    else
        echo "✅ Port $PORT is available"
    fi
fi

# Check PostgreSQL default port
if lsof -Pi :5432 -sTCP:LISTEN -t >/dev/null ; then
    echo "⚠️  Port 5432 (PostgreSQL) is already in use"
    echo "   This might conflict with Docker PostgreSQL. Consider stopping local PostgreSQL:"
    echo "   sudo systemctl stop postgresql"
else
    echo "✅ Port 5432 (PostgreSQL) is available"
fi

echo
echo "=== Summary ==="

if [ $missing_vars -eq 0 ]; then
    echo "🎉 Environment configuration looks good!"
    echo
    echo "Next steps:"
    echo "1. Start the application:"
    echo "   sudo docker compose up --build"
    echo
    echo "2. Or for development mode:"
    echo "   docker compose up postgres-zlagoda"
    echo "   ./migrate.sh"
    echo "   go run cmd/api/main.go"
    echo
    echo "3. Test the API:"
    echo "   curl http://localhost:$PORT/api/categories"
else
    echo "❌ Found $missing_vars configuration issue(s)"
    echo "   Please fix the issues above before starting the application"
    exit 1
fi

echo
echo "=== Helpful Commands ==="
echo "View logs:           docker compose logs"
echo "Connect to database: docker exec -it postgres-zlagoda psql -U $DB_USER -d $DB_NAME"
echo "Stop services:       docker compose down"
echo "Clean restart:       docker compose down && docker compose up --build"
