.PHONY: all clean build run stop test docker-build docker-run docker-stop docker-build-compose npm mysql docker-nuke

APP_NAME := main
NODE_MODULES := node_modules
BUILD := build
DOCKER_COMPOSE := docker-compose.yml

# Clean build artifacts
clean:
	rm -rf $(APP_NAME) $(NODE_MODULES) $(BUILD)

npm:
	npm install

# Build the application with docker-compose
build:
	docker-compose -f $(DOCKER_COMPOSE) build

# Run the application with docker-compose
run: npm build
	docker-compose -f $(DOCKER_COMPOSE) up

# Navigate througy mysql database
mysql:
	docker exec -it mysql-container bash

docker-nuke:
	docker stop `docker ps -qa`
	docker rm `docker ps -qa`
	docker rmi -f `docker images -qa`
	docker volume rm $(docker volume ls -qf)

# Stop the Docker containers
stop:
	docker-compose -f $(DOCKER_COMPOSE) down

# Run tests
test:
	go test -v ./...

# Run all targets
all: clean build test docker-run

# Help target
help:
	@echo "Available targets:"
	@echo "  build-local         - Build the application"
	@echo "  clean               - Clean build artifacts"
	@echo "  run-local           - Run the server locally"
	@echo "  run-live            - Run the server locally with live reloading using gin"
	@echo "  docker-build        - Build the Docker image"
	@echo "  docker-run          - Run the server using Docker"
	@echo "  docker-stop         - Stop the Docker containers"
	@echo "  build		      - Build the application with docker-compose"
	@echo "  run  		      - Run the application with docker-compose"
	@echo "  stop                - Stop the Docker containers"
	@echo "  test                - Run tests"
	@echo "  all                 - Clean, build, test, and run using Docker"
	@echo "  npm		      - installs node dependencies"
	@echo "  mysql		      - navigate through mysql database"
	@echo "  docker-nuke	      - nukes out all docker instances"
	@echo "  help                - Show this help message"

# Default target
default: help
