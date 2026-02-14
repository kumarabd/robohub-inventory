# RoboHub Inventory Service

![Docker Build](https://github.com/kumarabd/robohub-inventory/actions/workflows/docker-publish.yml/badge.svg)

A Golang microservice for managing inventory in the RoboHub robotics development platform. This service manages packages, repositories, scenarios, datasets, and simulators.

## Docker Images

Pre-built Docker images are available on GitHub Container Registry:

```bash
docker pull ghcr.io/kumarabd/robohub-inventory:latest
```

Replace `kumarabd` with your GitHub username or organization.

## Features

- HTTP server with graceful shutdown using chi router
- PostgreSQL database with GORM ORM
- **Automatic database migrations** with GORM AutoMigrate
- **Auto-detection of breaking schema changes**
- **Automatic seed data loading** for development
- Domain-driven design architecture
- RESTful API for inventory management
- Docker and docker-compose support
- Automated CI/CD with GitHub Actions
- Multi-platform Docker images (amd64, arm64)
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

### Quick Start

The service automatically handles database schema creation and migrations!

1. Start PostgreSQL:
```bash
docker run -d \
  --name robohub-postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=robohub_inventory \
  -p 5432:5432 \
  postgres:15-alpine
```

2. Run the service (migrations happen automatically):
```bash
# First run - creates tables and loads seed data
go run cmd/main.go

# Or with fresh data every time
FORCE_DROP_TABLES=true LOAD_SEED_DATA=true go run cmd/main.go
```

That's it! The service will:
- ✅ Create all tables automatically
- ✅ Apply schema changes
- ✅ Load sample data (if tables are empty)
- ✅ Detect and handle breaking changes

### Environment Variables for Migrations

- `FORCE_DROP_TABLES=true` - Drop and recreate all tables (⚠️ deletes data!)
- `LOAD_SEED_DATA=true` - Load sample data for testing/development

See [MIGRATION_GUIDE.md](MIGRATION_GUIDE.md) for details.

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

### Using Pre-built Docker Images

Pull and run the latest image from GitHub Container Registry:

```bash
# Pull the image
docker pull ghcr.io/kumarabd/robohub-inventory:latest

# Run with docker-compose or standalone
docker run -p 8080:8080 \
  -e DB_HOST=your-db-host \
  -e DB_USER=postgres \
  -e DB_PASSWORD=your-password \
  -e DB_NAME=robohub_inventory \
  ghcr.io/kumarabd/robohub-inventory:latest
```

## CI/CD Pipeline

This project includes automated CI/CD workflows using GitHub Actions:

- **Automatic builds**: Docker images are automatically built and published on every push to `main`/`develop` branches
- **Release tags**: Create a git tag (e.g., `v1.0.0`) to publish versioned releases
- **Multi-platform**: Images are built for both `linux/amd64` and `linux/arm64`
- **Manual builds**: Trigger custom builds via GitHub Actions UI

See [`.github/workflows/README.md`](.github/workflows/README.md) for detailed documentation.

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
