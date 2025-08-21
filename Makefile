# Docker services
LOCAL_DOCKER_COMPOSE_PROJECT_NAME=transactions_services_local
LOCAL_DOCKER_COMPOSE_FILE_PATH=./docker/local/docker-compose.yml

PROD_DOCKER_COMPOSE_PROJECT_NAME=transactions_services
PROD_DOCKER_COMPOSE_FILE_PATH=./docker/prod/docker-compose.yml

up_local_services:
	docker compose -p $(LOCAL_DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE_PATH) up -d
down_local_services:
	docker compose -p $(LOCAL_DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE_PATH) down
rebuild_local_services:
	docker compose -p $(LOCAL_DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE_PATH) up -d --build
restart_local_services: down_local_services up_local_services

up_prod_services:
	docker compose --env-file ./.env -p $(PROD_DOCKER_COMPOSE_PROJECT_NAME) -f $(PROD_DOCKER_COMPOSE_FILE_PATH) up -d
down_prod_services:
	docker compose --env-file ./.env -p $(PROD_DOCKER_COMPOSE_PROJECT_NAME) -f $(PROD_DOCKER_COMPOSE_FILE_PATH) down
rebuild_prod_services:
	docker compose --env-file ./.env -p $(PROD_DOCKER_COMPOSE_PROJECT_NAME) -f $(PROD_DOCKER_COMPOSE_FILE_PATH) up -d --build
restart_prod_services: down_prod_services up_prod_services

# Migrations(Goose)
MIGRATIONS_DIR=migrations
POSTGRES_DSN_DEFAULT=postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable

apply_migrations:
	goose -dir $(MIGRATIONS_DIR)  postgres "$(or $(POSTGRES_DSN), $(POSTGRES_DSN_DEFAULT))" up

add_migration:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# Deployment
build_prod_image:
	docker build -f .\docker\prod\Dockerfile .

deploy_prod:
	@echo "Deploying production services..."
	docker compose --env-file ./.env -p $(PROD_DOCKER_COMPOSE_PROJECT_NAME) -f $(PROD_DOCKER_COMPOSE_FILE_PATH) pull app
	docker compose --env-file ./.env -p $(PROD_DOCKER_COMPOSE_PROJECT_NAME) -f $(PROD_DOCKER_COMPOSE_FILE_PATH) up -d --force-recreate app
	@echo "Deployment complete."

