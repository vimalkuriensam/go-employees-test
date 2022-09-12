EMPLOYEES_BINARY=employeesApp
cwd  := $(shell pwd)

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -f ./deployment/docker-compose.yml -d --build
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_employees
	@echo "Stopping docker images (if running...)"
	docker-compose -f ./deployment/docker-compose.yml down
	@echo "Building (when required) and starting docker images..."
	docker-compose -f ./deployment/docker-compose.yml up --build -d
	@echo "Docker images built and started!"

build_employees:
	@echo "Building employees binary..."
	go build -o ${EMPLOYEES_BINARY} ./cmd/api/.
	@echo "Done!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose -f ./deployment/docker-compose.yml down
	@echo "Done!"