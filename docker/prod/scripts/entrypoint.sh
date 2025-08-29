#!/bin/sh
set -e

make apply_migrations

exec uvicorn main:app --host 0.0.0.0 --port "${INTERNAL_PORT}"
