#!/bin/bash

echo "Starting Zlagoda API service..."
echo "Environment variables:"
echo "  DB_HOST: $DB_HOST"
echo "  DB_PORT: $DB_PORT"
echo "  DB_NAME: $DB_NAME"
echo "  DB_USER: $DB_USER"
echo "  PORT: $PORT"
echo "  VAT_RATE: $VAT_RATE"

# Wait a moment to ensure database is ready
sleep 3

echo "Starting API server..."
exec /app/bin/api
