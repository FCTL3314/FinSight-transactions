#!/bin/sh
set -e

mkdir -p /app/logs
chown -R appuser:appuser /app/logs

gosu appuser uv run make apply_migrations

exec gosu appuser uv run uvicorn main:app --host 0.0.0.0 --port "${INTERNAL_PORT}"
