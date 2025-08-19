#!/bin/sh
set -e

make apply_migrations

exec ./app
