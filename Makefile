######    Variables    ######
# Environment
ENV_LOCAL_PATH=./settings/.env.local
ENV_PROD_PATH=./settings/.env.prod

# Docker
LOCAL_DOCKER_COMPOSE_PROJECT_NAME=transactions_services_local
LOCAL_DOCKER_COMPOSE_FILE_PATH=./docker/local/docker-compose.yml

PROD_DOCKER_COMPOSE_PROJECT_NAME=transactions_services
PROD_DOCKER_COMPOSE_FILE_PATH=./docker/prod/docker-compose.yml

# Migrations (Alembic)
ALEMBIC_CONFIG=./migrations/alembic.ini

######    Shortcuts    ######
# Local Docker Services
up_local_services:
	docker compose --env-file ${ENV_LOCAL_PATH} -p $(LOCAL_DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE_PATH) up -d
down_local_services:
	docker compose --env-file ${ENV_LOCAL_PATH} -p $(LOCAL_DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE_PATH) down
rebuild_local_services:
	docker compose --env-file ${ENV_LOCAL_PATH} -p $(LOCAL_DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE_PATH) up -d --build
restart_local_services: down_local_services up_local_services

# Production Docker Services
up_prod_services:
	docker compose --env-file ${ENV_PROD_PATH} -p $(PROD_DOCKER_COMPOSE_PROJECT_NAME) -f $(PROD_DOCKER_COMPOSE_FILE_PATH) up -d
down_prod_services:
	docker compose --env-file ${ENV_PROD_PATH} -p $(PROD_DOCKER_COMPOSE_PROJECT_NAME) -f $(PROD_DOCKER_COMPOSE_FILE_PATH) down
rebuild_prod_services:
	docker compose --env-file ${ENV_PROD_PATH} -p $(PROD_DOCKER_COMPOSE_PROJECT_NAME) -f $(PROD_DOCKER_COMPOSE_FILE_PATH) up -d --build
restart_prod_services: down_prod_services up_prod_services

# Migrations (Alembic)
apply_migrations:
	alembic -c $(ALEMBIC_CONFIG) upgrade head

create_migration:
	alembic -c $(ALEMBIC_CONFIG) revision --autogenerate -m "$(name)"

# Deployment
deploy_prod:
	@echo "Deploying production services..."
	docker compose --env-file ${ENV_PROD_PATH} -p $(PROD_DOCKER_COMPOSE_PROJECT_NAME) -f $(PROD_DOCKER_COMPOSE_FILE_PATH) pull app
	docker compose --env-file ${ENV_PROD_PATH} -p $(PROD_DOCKER_COMPOSE_PROJECT_NAME) -f $(PROD_DOCKER_COMPOSE_FILE_PATH) up -d --force-recreate app
	@echo "Deployment complete."

health_check_prod:
	docker compose --env-file ${ENV_PROD_PATH} -p $(PROD_DOCKER_COMPOSE_PROJECT_NAME) -f $(PROD_DOCKER_COMPOSE_FILE_PATH) ps app
