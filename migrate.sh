#!/bin/bash

echo "Running migrations"

# Wait for database to be ready
echo "Waiting for database to be ready..."
for i in {1..30}; do
    if migrate -database "${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path db/migrations version &>/dev/null; then
        echo "Database is ready!"
        break
    fi
    echo "Database not ready, waiting... ($i/30)"
    sleep 2
done

# Run migrations
echo "Running database migrations..."
migrate -database "${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path db/migrations up

if [ $? -eq 0 ]; then
    echo "Migrations completed successfully!"
else
    echo "Migration failed!"
    exit 1
fi
