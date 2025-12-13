#!/bin/bash
set -e

echo "Running database migrations..."
goose -dir ./sql/schema postgres "$DATABASE_URL" up

echo "Starting server..."
exec ./server
