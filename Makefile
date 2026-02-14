.PHONY: help build run test clean docker-build docker-run docker-stop docker-compose-up docker-compose-down migrate

# Variables
APP_NAME=robohub-inventory
DOCKER_IMAGE=$(APP_NAME):latest
PORT=8080

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the Go application
	@echo "Building $(APP_NAME)..."
	@go build -o bin/$(APP_NAME) cmd/main.go

run: ## Run the application locally
	@echo "Running $(APP_NAME) on port $(PORT)..."
	@PORT=$(PORT) ./bin/$(APP_NAME)

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean

docker-build: ## Build Docker image
	@echo "Building Docker image $(DOCKER_IMAGE)..."
	@docker build -t $(DOCKER_IMAGE) .

docker-run: ## Run Docker container
	@echo "Running Docker container on port $(PORT)..."
	@docker run -d -p $(PORT):8080 --name $(APP_NAME) $(DOCKER_IMAGE)

docker-stop: ## Stop and remove Docker container
	@echo "Stopping Docker container..."
	@docker stop $(APP_NAME) || true
	@docker rm $(APP_NAME) || true

docker-logs: ## Show Docker container logs
	@docker logs -f $(APP_NAME)

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

docker-compose-up: ## Start docker-compose services (app + postgres)
	@echo "Starting docker-compose services..."
	@docker-compose up -d

docker-compose-down: ## Stop docker-compose services
	@echo "Stopping docker-compose services..."
	@docker-compose down

migrate: ## Run database migrations (requires DB connection)
	@echo "Running database migrations..."
	@go run main.go migrate || echo "Note: Migrations run automatically on startup"
