include scripts/*.mk

DEV_COMPOSE_ARGS=--env-file .env.dev -f docker-compose.dev.yaml
DEV_COMPOSE=docker compose $(DEV_COMPOSE_ARGS)

dev-build:
	$(DEV_COMPOSE) build --no-cache

dev-up: api_docker_build consumer_docker_build
	$(DEV_COMPOSE) up -d --no-deps --build

dev-down:
	$(DEV_COMPOSE) down

dev-logs:
	$(DEV_COMPOSE) logs

dev-migrate-up:
	$(DEV_COMPOSE) up -d migrations-up

dev-migrate-down:
	$(DEV_COMPOSE) --profile migrations-down up -d migrations-down

dev-api-run:
	go run cmd/api/api.go
	
dev-api-up:
	$(DEV_COMPOSE) --profile api up -d api

build-kafka:
	docker build -t kafka-jmx -f Dockerfile.kafka .

build-zookeeper:
	docker build -t zookeeper-jmx -f Dockerfile.zookeeper .

dev-db:
	PATH=/opt/homebrew/bin/:PATH psql "postgres://dev:dev1234@localhost:5432/dev?sslmode=disable"
redis-connect:
	docker exec -it magic_potions.redis redis-cli -p 6379
