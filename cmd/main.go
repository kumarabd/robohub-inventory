package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"robohub-inventory/internal/config"
	"robohub-inventory/internal/database"
	"robohub-inventory/internal/http"
	"robohub-inventory/internal/logger"
	"robohub-inventory/internal/metrics"
	"robohub-inventory/pkg/dataset"
	pkg "robohub-inventory/pkg/package"
	"robohub-inventory/pkg/repository"
	"robohub-inventory/pkg/scenario"
	"robohub-inventory/pkg/simulator"
)

func main() {
	// Initialize logger
	log := logger.New()
	log.Info("Starting RoboHub Inventory Service...")

	// Initialize metrics
	metrics.Init()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration: %v", err)
	}

	// Connect to database
	db, err := database.Connect(&cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to run migrations: %v", err)
	}
	log.Info("Database migrations completed")

	// Initialize repositories
	pkgRepo := pkg.NewRepository(db)
	repoRepo := repository.NewRepository(db)
	scenarioRepo := scenario.NewRepository(db)
	datasetRepo := dataset.NewRepository(db)
	simulatorRepo := simulator.NewRepository(db)

	// Initialize services
	pkgService := pkg.NewService(pkgRepo)
	repoService := repository.NewService(repoRepo)
	scenarioService := scenario.NewService(scenarioRepo)
	datasetService := dataset.NewService(datasetRepo)
	simulatorService := simulator.NewService(simulatorRepo)

	// Initialize router
	router := http.NewRouter(
		pkgService,
		repoService,
		scenarioService,
		datasetService,
		simulatorService,
	)

	// Initialize HTTP server
	server := http.NewServer(&cfg.Server, router)

	// Start server in a goroutine
	go func() {
		log.Info("Server starting on %s:%s", cfg.Server.Host, cfg.Server.Port)
		if err := server.Start(); err != nil && err != context.Canceled {
			log.Fatal("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: %v", err)
	}

	log.Info("Server exited")
}
