#!/bin/sh
set -e

# chown -R appuser:appuser /app/logs

make apply_migrations

exec uvicorn main:app --host 0.0.0.0 --port "${INTERNAL_PORT}"
