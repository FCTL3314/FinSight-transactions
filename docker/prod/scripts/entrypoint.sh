#!/bin/sh
set -e

chown -R appuser:appuser /app/logs

make apply_migrations POSTGRES_DSN=postgresql://postgres:postgres@db:5432/postgres?sslmode=disable

exec su-exec appuser ./app