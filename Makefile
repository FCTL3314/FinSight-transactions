# Migrations(Goose)
MIGRATIONS_DIR=migrations
POSTGRES_DSN_DEFAULT=postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable

apply_migrations:
	goose -dir $(MIGRATIONS_DIR)  postgres "$(or $(POSTGRES_DSN), $(POSTGRES_DSN_DEFAULT))" up

add_migration:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql