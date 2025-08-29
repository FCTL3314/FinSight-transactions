#!/bin/sh
set -e

uv run make apply_migrations

uv run uvicorn main:app --host 0.0.0.0 --port "${INTERNAL_PORT}"
