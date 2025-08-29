#!/bin/sh
set -e

uv run uvicorn main:app --host 0.0.0.0 --port "${INTERNAL_PORT}" --reload
