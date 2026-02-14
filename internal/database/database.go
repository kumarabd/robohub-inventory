package database

import (
	"fmt"
	"robohub-inventory/internal/config"
	"robohub-inventory/pkg/dataset"
	"robohub-inventory/pkg/package"
	"robohub-inventory/pkg/repository"
	"robohub-inventory/pkg/scenario"
	"robohub-inventory/pkg/simulator"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the global database instance
var DB *gorm.DB

// Connect initializes the database connection
func Connect(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	return db, nil
}

// Migrate runs database migrations
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&pkg.Package{},
		&repository.Repository{},
		&scenario.Scenario{},
		&dataset.Dataset{},
		&simulator.Simulator{},
	)
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
