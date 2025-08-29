#!/bin/sh
set -e

# chown -R appuser:appuser /app/logs

alembic -c settings/alembic.ini upgrade head

exec uvicorn main:app --host 0.0.0.0 --port "${INTERNAL_PORT}"
