PROJECT?=github.com/donskova1ex/magic_potions
CONSUMER_NAME?=recipes-consumer
CONSUMER_VERSION?=0.0.1
CONSUMER_CONTAINER_NAME?=docker.io/donskova1ex/${CONSUMER_NAME}

# for local uses
consumer_local_build:
	go build -o bin/${CONSUMER_NAME} cmd/${CONSUMER_NAME}/main.go

# docker
consumer_docker_build:
	docker build -t ${CONSUMER_CONTAINER_NAME}:${API_VERSION} -t ${CONSUMER_CONTAINER_NAME}:latest -f Dockerfile.consumer .
consumer_docker_push:
	docker push ${CONSUMER_CONTAINER_NAME}:${CONSUMER_VERSION}
	docker push ${CONSUMER_CONTAINER_NAME}:latest
consumer_docker_update: consumer_docker_build consumer_docker_push