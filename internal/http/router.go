package http

import (
	"encoding/json"
	"net/http"

	"robohub-inventory/internal/http/handlers"
	"robohub-inventory/pkg/dataset"
	pkg "robohub-inventory/pkg/package"
	"robohub-inventory/pkg/repository"
	"robohub-inventory/pkg/scenario"
	"robohub-inventory/pkg/simulator"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(
	pkgService *pkg.Service,
	repoService *repository.Service,
	scenarioService *scenario.Service,
	datasetService *dataset.Service,
	simulatorService *simulator.Service,
) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Handlers
	healthHandler := handlers.NewHealthHandler()
	packageHandler := handlers.NewPackageHandler(pkgService)
	repositoryHandler := handlers.NewRepositoryHandler(repoService)
	scenarioHandler := handlers.NewScenarioHandler(scenarioService)
	datasetHandler := handlers.NewDatasetHandler(datasetService)
	simulatorHandler := handlers.NewSimulatorHandler(simulatorService)

	// Routes
	r.Get("/health", healthHandler.Health)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Welcome to RoboHub Inventory Service",
			"version": "1.0.0",
		})
	})

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		// Packages
		r.Route("/packages", func(r chi.Router) {
			r.Post("/", packageHandler.CreatePackage)
			r.Get("/", packageHandler.ListPackages)
			r.Get("/{id}", packageHandler.GetPackage)
			r.Put("/{id}", packageHandler.UpdatePackage)
			r.Delete("/{id}", packageHandler.DeletePackage)
		})

		// Repositories
		r.Route("/repositories", func(r chi.Router) {
			r.Post("/", repositoryHandler.CreateRepository)
			r.Get("/", repositoryHandler.ListRepositories)
			r.Get("/{id}", repositoryHandler.GetRepository)
			r.Put("/{id}", repositoryHandler.UpdateRepository)
			r.Delete("/{id}", repositoryHandler.DeleteRepository)
		})

		// Scenarios
		r.Route("/scenarios", func(r chi.Router) {
			r.Post("/", scenarioHandler.CreateScenario)
			r.Get("/", scenarioHandler.ListScenarios)
			r.Get("/{id}", scenarioHandler.GetScenario)
			r.Put("/{id}", scenarioHandler.UpdateScenario)
			r.Delete("/{id}", scenarioHandler.DeleteScenario)
		})

		// Datasets
		r.Route("/datasets", func(r chi.Router) {
			r.Post("/", datasetHandler.CreateDataset)
			r.Get("/", datasetHandler.ListDatasets)
			r.Get("/{id}", datasetHandler.GetDataset)
			r.Put("/{id}", datasetHandler.UpdateDataset)
			r.Delete("/{id}", datasetHandler.DeleteDataset)
		})

		// Simulators
		r.Route("/simulators", func(r chi.Router) {
			r.Post("/", simulatorHandler.CreateSimulator)
			r.Get("/", simulatorHandler.ListSimulators)
			r.Get("/{id}", simulatorHandler.GetSimulator)
			r.Put("/{id}", simulatorHandler.UpdateSimulator)
			r.Delete("/{id}", simulatorHandler.DeleteSimulator)
		})
	})

	return r
}
