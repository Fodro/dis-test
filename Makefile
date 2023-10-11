MAKEFLAGS += --silent

export ENV_FILE = deployments/.env

ifneq ("$(wildcard $(ENV_FILE))","")
	include $(ENV_FILE)
	export $(shell sed 's/=.*//' $(ENV_FILE))
endif

export DOCKER_COMPOSE = docker-compose -f deployments/docker-compose.yml
DATABASE_URL = postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@127.0.0.1:$(POSTGRES_PORT)/$(COMPOSE_PROJECT_NAME)?sslmode=disable

.PHONY: *

# Installation
install:
	./scripts/install.sh
init: start
	./scripts/init.sh
reinstall: clean install init
clean:
	$(DOCKER_COMPOSE) down --rmi local -v
rebuild:
	$(DOCKER_COMPOSE) up -d --no-deps --build api

start:
	$(DOCKER_COMPOSE) up -d
restart: stop start
stop:
	$(DOCKER_COMPOSE) down

test:
	$(DOCKER_COMPOSE) exec -T api go test -v -cover ./...

lint-api:
	docker run --rm -v $(PWD)/:/app -w /server golangci/golangci-lint:v1.31.0 golangci-lint run -v

cs-fix-api:
	$(DOCKER_COMPOSE) exec api go fmt ./...

# Entering into containers
enter-nginx:
	$(DOCKER_COMPOSE) exec nginx bash
enter-api:
	$(DOCKER_COMPOSE) exec api bash
enter-postgres:
	$(DOCKER_COMPOSE) exec postgres bash

# Showing logs
logs-nginx:
	$(DOCKER_COMPOSE) logs nginx
logs-api:
	$(DOCKER_COMPOSE) logs api
logs-postgres:
	$(DOCKER_COMPOSE) logs postgres
logs-kafka:
	$(DOCKER_COMPOSE) logs kafka
logs-zookeeper:
	$(DOCKER_COMPOSE) logs zookeeper

ps:
	$(DOCKER_COMPOSE) ps

generate-openapi:
	$(DOCKER_COMPOSE) exec -T api swag init -g cmd/discipline-api/main.go --propertyStrategy snakecase --parseDependency

migrate-up:
	echo "It will be migrated soon..."

# Global
.DEFAULT_GOAL := install
