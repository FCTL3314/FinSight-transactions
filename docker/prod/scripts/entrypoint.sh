#!/bin/sh
set -e

chown -R appuser:appuser /app/logs

gosu appuser make apply_migrations POSTGRES_DSN=postgresql://postgres:postgres@db:5432/postgres?sslmode=disable

exec gosu appuser ./app