#!/bin/sh
set -e

mkdir -p /app/logs
chown -R appuser:appuser /app/logs

make apply_migrations POSTGRES_DSN=postgresql://postgres:postgres@db:5432/postgres?sslmode=disable

exec ./app
