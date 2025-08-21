#!/bin/sh
set -e

make apply_migrations POSTGRES_DSN=postgresql://postgres:postgres@db:5432/postgres?sslmode=disable

exec ./app
