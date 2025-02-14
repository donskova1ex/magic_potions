include scripts/*.mk

DEV_COMPOSE_ARGS=--env-file .env.dev -f docker-compose.dev.yaml
DEV_COMPOSE_ENV=docker compose $(DEV_COMPOSE_ARGS)
DEV_COMPOSE=docker compose $(DEV_COMPOSE_ARGS)

dev-build:
	$(DEV_COMPOSE) build --no-cache

dev-up: api_docker_build consumer_docker_build dev-build
	$(DEV_COMPOSE) up -d --no-deps --build

dev-up-env: dev-build
	$(DEV_COMPOSE_ENV) up -d 

dev-down:
	$(DEV_COMPOSE) down

dev-logs:
	docker compose -f docker-compose.dev.yaml logs

dev-migrate-up:
	$(DEV_COMPOSE) up -d migrations-up

dev-migrate-down:
	$(DEV_COMPOSE) --profile migrations-down up -d migrations-down

dev-api-run:
	go run cmd/api/api.go
	
build-kafka:
	docker build -t kafka-jmx -f Dockerfile.kafka .

build-zookeeper:
	docker build -t zookeeper-jmx -f Dockerfile.zookeeper .

dev-db-login:
	psql "postgres://dev:dev1234@localhost:5432/dev?sslmode=disable"