# RoboHub Inventory Service

A Golang microservice for managing inventory in the RoboHub robotics development platform. This service manages packages, repositories, scenarios, datasets, and simulators.

## Features

- HTTP server with graceful shutdown using chi router
- PostgreSQL database with GORM ORM
- Domain-driven design architecture
- RESTful API for inventory management
- Docker and docker-compose support
- Health check endpoint
- Structured logging
- Metrics collection

## Architecture

The project follows Domain-Driven Design principles with the following structure:

```
.
├── pkg/                    # Domain packages (public)
│   ├── package/           # Package entity, repository, service
│   ├── repository/        # Repository entity, repository, service
│   ├── scenario/          # Scenario entity, repository, service
│   ├── dataset/           # Dataset entity, repository, service
│   └── simulator/         # Simulator entity, repository, service
├── internal/              # Internal packages (private)
│   ├── config/           # Configuration management
│   ├── database/         # Database connection and migrations
│   ├── http/             # HTTP server, router, handlers
│   ├── logger/           # Logging utilities
│   └── metrics/          # Metrics collection
└── main.go               # Application entry point
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 15 or higher
- Docker and Docker Compose (optional, for containerized deployment)
- Make (optional, for using Makefile commands)

## Getting Started

### Local Development

1. Start PostgreSQL (using docker-compose):
```bash
make docker-compose-up
```

Or manually:
```bash
docker run -d \
  --name robohub-postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=robohub_inventory \
  -p 5432:5432 \
  postgres:15-alpine
```

2. Install dependencies:
```bash
make deps
```

3. Set environment variables (optional, defaults provided):
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=robohub_inventory
export PORT=8080
```

4. Build the application:
```bash
make build
```

5. Run the application:
```bash
make run
```

### Using Docker Compose

1. Start all services (app + postgres):
```bash
make docker-compose-up
```

2. Stop all services:
```bash
make docker-compose-down
```

3. View logs:
```bash
docker-compose logs -f app
```

## API Endpoints

### Health & Info
- `GET /` - Root endpoint with service information
- `GET /health` - Health check endpoint

### Packages
- `POST /api/v1/packages` - Create a new package
- `GET /api/v1/packages` - List packages (query params: `limit`, `offset`)
- `GET /api/v1/packages/{id}` - Get package by ID
- `PUT /api/v1/packages/{id}` - Update package
- `DELETE /api/v1/packages/{id}` - Delete package

### Repositories
- `POST /api/v1/repositories` - Create a new repository
- `GET /api/v1/repositories` - List repositories (query params: `limit`, `offset`)
- `GET /api/v1/repositories/{id}` - Get repository by ID
- `PUT /api/v1/repositories/{id}` - Update repository
- `DELETE /api/v1/repositories/{id}` - Delete repository

### Scenarios
- `POST /api/v1/scenarios` - Create a new scenario
- `GET /api/v1/scenarios` - List scenarios (query params: `limit`, `offset`)
- `GET /api/v1/scenarios/{id}` - Get scenario by ID
- `PUT /api/v1/scenarios/{id}` - Update scenario
- `DELETE /api/v1/scenarios/{id}` - Delete scenario

### Datasets
- `POST /api/v1/datasets` - Create a new dataset
- `GET /api/v1/datasets` - List datasets (query params: `limit`, `offset`)
- `GET /api/v1/datasets/{id}` - Get dataset by ID
- `PUT /api/v1/datasets/{id}` - Update dataset
- `DELETE /api/v1/datasets/{id}` - Delete dataset

### Simulators
- `POST /api/v1/simulators` - Create a new simulator
- `GET /api/v1/simulators` - List simulators (query params: `limit`, `offset`)
- `GET /api/v1/simulators/{id}` - Get simulator by ID
- `PUT /api/v1/simulators/{id}` - Update simulator
- `DELETE /api/v1/simulators/{id}` - Delete simulator

## Environment Variables

- `PORT` - Server port (default: 8080)
- `HOST` - Server host (default: 0.0.0.0)
- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `DB_USER` - Database user (default: postgres)
- `DB_PASSWORD` - Database password (default: postgres)
- `DB_NAME` - Database name (default: robohub_inventory)
- `DB_SSLMODE` - SSL mode (default: disable)

## Makefile Commands

Run `make help` to see all available commands:

- `make build` - Build the application
- `make run` - Run the application locally
- `make test` - Run tests
- `make clean` - Clean build artifacts
- `make docker-build` - Build Docker image
- `make docker-run` - Run Docker container
- `make docker-stop` - Stop Docker container
- `make docker-logs` - Show Docker logs
- `make docker-compose-up` - Start docker-compose services
- `make docker-compose-down` - Stop docker-compose services
- `make deps` - Download dependencies
- `make fmt` - Format code
- `make vet` - Run go vet
- `make migrate` - Run database migrations (runs automatically on startup)

## Example API Usage

### Create a Package
```bash
curl -X POST http://localhost:8080/api/v1/packages \
  -H "Content-Type: application/json" \
  -d '{
    "name": "ros-navigation",
    "version": "1.0.0",
    "description": "ROS Navigation Stack",
    "type": "ros",
    "repository": "https://github.com/ros-planning/navigation2",
    "tags": ["ros", "navigation", "robotics"]
  }'
```

### List Packages
```bash
curl http://localhost:8080/api/v1/packages?limit=10&offset=0
```

### Get Package by ID
```bash
curl http://localhost:8080/api/v1/packages/1
```

## Technology Stack

- **Go 1.21** - Programming language
- **Chi Router** - HTTP router and middleware
- **GORM** - ORM library
- **PostgreSQL** - Database
- **Docker** - Containerization

## License

MIT
