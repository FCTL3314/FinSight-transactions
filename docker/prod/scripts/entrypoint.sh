#!/bin/sh
set -e

chown -R appuser:appuser /app/logs

uv run make apply_migrations

exec su-exec appuser uv run uvicorn main:app --host 0.0.0.0 --port "${INTERNAL_PORT}"
