package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"robohub-inventory/internal/config"
	"robohub-inventory/pkg/dataset"
	pkg "robohub-inventory/pkg/package"
	"robohub-inventory/pkg/repository"
	"robohub-inventory/pkg/scenario"
	"robohub-inventory/pkg/simulator"
)

// DB is the global database instance
var DB *gorm.DB

// Connect initializes the database connection and runs migrations
func Connect(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db

	// Run migrations automatically
	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database connected and migrated successfully")
	return db, nil
}

// runMigrations handles automatic schema migrations
func runMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")

	// Check if we need to drop tables (breaking changes)
	if shouldDropTables(db) {
		log.Println("Detected schema breaking changes, dropping existing tables...")
		if err := dropAllTables(db); err != nil {
			return fmt.Errorf("failed to drop tables: %w", err)
		}
	}

	// Enable UUID extension for PostgreSQL
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"pgcrypto\"").Error; err != nil {
		log.Printf("Warning: Could not enable pgcrypto extension: %v", err)
	}

	// AutoMigrate all models
	if err := Migrate(db); err != nil {
		return err
	}

	log.Println("Migrations completed successfully")

	// Load seed data if tables are empty
	if shouldLoadSeedData(db) {
		log.Println("Loading seed data...")
		if err := loadSeedData(db); err != nil {
			log.Printf("Warning: Failed to load seed data: %v", err)
		}
	}

	return nil
}

// Migrate runs database migrations using GORM AutoMigrate
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&repository.Repository{},
		&pkg.Package{},
		&scenario.Scenario{},
		&dataset.Dataset{},
		&simulator.Simulator{},
	)
}

// shouldDropTables checks if we need to drop tables due to breaking changes
func shouldDropTables(db *gorm.DB) bool {
	// Check if FORCE_DROP environment variable is set
	if os.Getenv("FORCE_DROP_TABLES") == "true" {
		return true
	}

	// Check if tables exist with old schema (uint IDs instead of UUIDs)
	var count int64
	
	// Check if repositories table exists and has integer ID
	result := db.Raw(`
		SELECT COUNT(*) 
		FROM information_schema.columns 
		WHERE table_name = 'repositories' 
		AND column_name = 'id' 
		AND data_type IN ('integer', 'bigint')
	`).Scan(&count)
	
	if result.Error == nil && count > 0 {
		log.Println("Detected old schema with integer IDs - will drop and recreate tables")
		return true
	}

	return false
}

// dropAllTables drops all application tables
func dropAllTables(db *gorm.DB) error {
	// Drop in reverse order to handle foreign keys
	tables := []string{
		"simulators",
		"datasets",
		"scenarios",
		"packages",
		"repositories",
	}

	for _, table := range tables {
		if err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", table)).Error; err != nil {
			return fmt.Errorf("failed to drop table %s: %w", table, err)
		}
		log.Printf("Dropped table: %s", table)
	}

	return nil
}

// shouldLoadSeedData checks if seed data should be loaded
func shouldLoadSeedData(db *gorm.DB) bool {
	// Only load if LOAD_SEED_DATA env var is set or tables are empty
	if os.Getenv("LOAD_SEED_DATA") == "true" {
		return true
	}

	var count int64
	db.Model(&repository.Repository{}).Count(&count)
	return count == 0
}

// loadSeedData loads sample data into the database
func loadSeedData(db *gorm.DB) error {
	// Create sample repositories
	repos := []*repository.Repository{
		{
			Name:          "ros-planning/navigation2",
			Provider:      "github",
			URL:           "https://github.com/ros-planning/navigation2",
			Description:   "ROS 2 Navigation Stack",
			DefaultBranch: "main",
			Visibility:    "public",
			SyncStatus:    "synced",
			AutoSync:      true,
			LatestCommit: repository.LatestCommit{
				Hash:    "a1b2c3d4e5f6",
				Message: "Add new planner plugin",
				Author:  "John Doe",
				Date:    time.Now().Add(-24 * time.Hour),
				URL:     "https://github.com/ros-planning/navigation2/commit/a1b2c3d4e5f6",
			},
			WebhookStatus: "active",
			Tags:          []string{"ros2", "navigation", "autonomous"},
			PackageCount:  0,
			Owner: repository.Owner{
				ID:        "user-001",
				Name:      "ros-planning",
				AvatarURL: "https://avatars.githubusercontent.com/ros-planning",
			},
		},
		{
			Name:          "ros-perception/perception_pcl",
			Provider:      "github",
			URL:           "https://github.com/ros-perception/perception_pcl",
			Description:   "PCL (Point Cloud Library) ROS interface",
			DefaultBranch: "ros2",
			Visibility:    "public",
			SyncStatus:    "synced",
			AutoSync:      true,
			LatestCommit: repository.LatestCommit{
				Hash:    "b2c3d4e5f6a1",
				Message: "Update point cloud filters",
				Author:  "Jane Smith",
				Date:    time.Now().Add(-48 * time.Hour),
				URL:     "https://github.com/ros-perception/perception_pcl/commit/b2c3d4e5f6a1",
			},
			WebhookStatus: "active",
			Tags:          []string{"ros2", "perception", "point-cloud"},
			PackageCount:  0,
			Owner: repository.Owner{
				ID:        "user-002",
				Name:      "ros-perception",
				AvatarURL: "https://avatars.githubusercontent.com/ros-perception",
			},
		},
	}

	for _, repo := range repos {
		if err := db.Create(repo).Error; err != nil {
			return fmt.Errorf("failed to create repository: %w", err)
		}
	}

	// Create sample packages
	packages := []*pkg.Package{
		{
			Name:          "nav2_planner",
			DisplayName:   "Nav2 Planner",
			Description:   "Global path planning server for Nav2",
			RepoID:        repos[0].ID,
			RepoName:      repos[0].Name,
			Path:          "nav2_planner",
			Types:         []string{"planner", "navigation"},
			LatestVersion: "1.1.9",
			Versions:      []string{"1.1.9", "1.1.8", "1.1.7"},
			Tags:          []string{"navigation", "planning", "ros2"},
			Keywords:      []string{"path-planning", "global-planner", "navigation"},
			ValidationStatus: pkg.ValidationStatus{
				LastValidated: time.Now().Add(-1 * time.Hour),
				Status:        "pass",
				PassRate:      95.5,
			},
			Owner: pkg.Owner{
				ID:        "user-001",
				Name:      "ros-planning",
				AvatarURL: "https://avatars.githubusercontent.com/ros-planning",
			},
		},
		{
			Name:          "nav2_controller",
			DisplayName:   "Nav2 Controller",
			Description:   "Local trajectory planning and control for Nav2",
			RepoID:        repos[0].ID,
			RepoName:      repos[0].Name,
			Path:          "nav2_controller",
			Types:         []string{"control", "navigation"},
			LatestVersion: "1.1.9",
			Versions:      []string{"1.1.9", "1.1.8", "1.1.7"},
			Tags:          []string{"navigation", "control", "ros2"},
			Keywords:      []string{"trajectory", "controller", "dwa"},
			ValidationStatus: pkg.ValidationStatus{
				LastValidated: time.Now().Add(-2 * time.Hour),
				Status:        "pass",
				PassRate:      92.3,
			},
			Owner: pkg.Owner{
				ID:        "user-001",
				Name:      "ros-planning",
				AvatarURL: "https://avatars.githubusercontent.com/ros-planning",
			},
		},
		{
			Name:          "pcl_ros",
			DisplayName:   "PCL ROS",
			Description:   "Point Cloud Library ROS2 integration",
			RepoID:        repos[1].ID,
			RepoName:      repos[1].Name,
			Path:          "pcl_ros",
			Types:         []string{"perception", "sensors"},
			LatestVersion: "2.5.0",
			Versions:      []string{"2.5.0", "2.4.0", "2.3.0"},
			Tags:          []string{"perception", "point-cloud", "ros2"},
			Keywords:      []string{"pcl", "3d-vision", "lidar"},
			ValidationStatus: pkg.ValidationStatus{
				LastValidated: time.Now().Add(-3 * time.Hour),
				Status:        "pass",
				PassRate:      88.7,
			},
			Owner: pkg.Owner{
				ID:        "user-002",
				Name:      "ros-perception",
				AvatarURL: "https://avatars.githubusercontent.com/ros-perception",
			},
		},
	}

	for _, p := range packages {
		if err := db.Create(p).Error; err != nil {
			return fmt.Errorf("failed to create package: %w", err)
		}
	}

	// Create sample scenarios
	scenarios := []*scenario.Scenario{
		{
			Name:                "Warehouse Navigation Basic",
			Slug:                "warehouse-nav-basic",
			Description:         "Navigate through a basic warehouse environment with static obstacles",
			Category:            "navigation",
			Difficulty:          "easy",
			MaintainedBy:        "RoboHub",
			Verified:            true,
			WhatItTests:         []string{"Obstacle avoidance", "Path planning", "Goal reaching"},
			WhyItMatters:        "Validates basic navigation capabilities in structured environments",
			RealWorldAnalogs:    []string{"Amazon fulfillment center", "Retail warehouse"},
			Domain:              "indoor",
			SupportedSimulators: []string{"Gazebo", "CARLA", "Unity"},
			RequiredInputs: []scenario.RequiredInput{
				{Name: "start_pose", Type: "geometry_msgs/PoseStamped", Description: "Starting position"},
				{Name: "goal_pose", Type: "geometry_msgs/PoseStamped", Description: "Target position"},
			},
			SuccessCriteria: []scenario.SuccessCriterion{
				{Name: "Success Rate", Description: "Percentage of successful goal reaches", Threshold: ">90%", Unit: "percentage"},
				{Name: "Path Efficiency", Description: "Path length vs optimal", Threshold: "<120%", Unit: "percentage"},
			},
			PassDefinition: "Robot reaches goal without collisions within time limit",
			Tags:           []string{"navigation", "warehouse", "basic"},
			Owner: scenario.Owner{
				ID:   "robohub-001",
				Name: "RoboHub Team",
			},
			Version: "1.0.0",
		},
		{
			Name:                "Urban Autonomous Driving",
			Slug:                "urban-autonomous-driving",
			Description:         "Navigate through urban environment with dynamic obstacles and traffic rules",
			Category:            "navigation",
			Difficulty:          "hard",
			MaintainedBy:        "Community",
			Verified:            true,
			WhatItTests:         []string{"Dynamic obstacle avoidance", "Traffic rule compliance", "Lane keeping"},
			WhyItMatters:        "Tests autonomous vehicle capabilities in complex real-world scenarios",
			RealWorldAnalogs:    []string{"City streets", "Downtown traffic"},
			Domain:              "urban",
			SupportedSimulators: []string{"CARLA", "AirSim"},
			RequiredInputs: []scenario.RequiredInput{
				{Name: "route", Type: "nav_msgs/Path", Description: "Planned route"},
				{Name: "traffic_rules", Type: "json", Description: "Local traffic regulations"},
			},
			SuccessCriteria: []scenario.SuccessCriterion{
				{Name: "Safety Score", Description: "No collisions or violations", Threshold: "100%", Unit: "percentage"},
				{Name: "Arrival Time", Description: "Within expected time window", Threshold: "±10%", Unit: "percentage"},
			},
			PassDefinition: "Complete route safely while following all traffic rules",
			Tags:           []string{"autonomous-driving", "urban", "advanced"},
			Owner: scenario.Owner{
				ID:   "community-001",
				Name: "AV Community",
			},
			Version: "2.1.0",
		},
		{
			Name:                "Object Detection Indoor",
			Slug:                "object-detection-indoor",
			Description:         "Detect and classify objects in indoor environment using camera and lidar",
			Category:            "perception",
			Difficulty:          "medium",
			MaintainedBy:        "Partner",
			Verified:            true,
			WhatItTests:         []string{"Object detection accuracy", "Classification performance", "Multi-sensor fusion"},
			WhyItMatters:        "Validates perception pipeline for indoor manipulation tasks",
			RealWorldAnalogs:    []string{"Home assistance", "Office automation"},
			Domain:              "indoor",
			SupportedSimulators: []string{"Gazebo", "Webots"},
			RequiredInputs: []scenario.RequiredInput{
				{Name: "sensor_data", Type: "sensor_msgs/PointCloud2", Description: "3D sensor data"},
				{Name: "camera_image", Type: "sensor_msgs/Image", Description: "RGB camera feed"},
			},
			SuccessCriteria: []scenario.SuccessCriterion{
				{Name: "Detection Rate", Description: "Percentage of objects detected", Threshold: ">85%", Unit: "percentage"},
				{Name: "False Positives", Description: "Incorrect detections", Threshold: "<5%", Unit: "percentage"},
			},
			PassDefinition: "Detect at least 85% of objects with less than 5% false positives",
			Tags:           []string{"perception", "object-detection", "indoor"},
			Owner: scenario.Owner{
				ID:   "partner-001",
				Name: "TechPartner Inc",
			},
			Version: "1.5.0",
		},
	}

	for _, s := range scenarios {
		if err := db.Create(s).Error; err != nil {
			return fmt.Errorf("failed to create scenario: %w", err)
		}
	}

	// Create sample datasets
	datasets := []*dataset.Dataset{
		{
			Name:         "Warehouse Navigation Dataset v1",
			Slug:         "warehouse-nav-v1",
			Description:  "Indoor warehouse navigation data with lidar and camera feeds",
			Type:         "robotics",
			Modality:     "multimodal",
			Format:       "rosbag2",
			License:      "MIT",
			Tags:         []string{"warehouse", "navigation", "indoor"},
			WhatsInside:  []string{"Lidar scans", "RGB camera images", "Odometry", "Ground truth poses"},
			SizeGB:       15.5,
			SamplesCount: 10000,
			Duration:     3600,
			Source:       "uploaded",
			OwnerType:    "organization",
			OwnerID:      "org-001",
			OwnerName:    "RoboHub Labs",
			Visibility:   "public",
		},
		{
			Name:         "Urban Driving CARLA",
			Slug:         "urban-driving-carla",
			Description:  "Synthetic urban driving data generated in CARLA simulator",
			Type:         "autonomous-driving",
			Modality:     "multimodal",
			Format:       "parquet",
			License:      "CC-BY",
			Tags:         []string{"autonomous-driving", "urban", "synthetic"},
			WhatsInside:  []string{"RGB cameras", "Depth images", "Semantic segmentation", "Vehicle telemetry"},
			SizeGB:       50.2,
			SamplesCount: 25000,
			Duration:     7200,
			Source:       "partner",
			OwnerType:    "organization",
			OwnerID:      "org-002",
			OwnerName:    "CARLA Team",
			Visibility:   "public",
		},
		{
			Name:         "Indoor Object Recognition",
			Slug:         "indoor-object-recognition",
			Description:  "Labeled indoor objects dataset for perception tasks",
			Type:         "indoor-mapping",
			Modality:     "camera",
			Format:       "hdf5",
			License:      "Apache-2.0",
			Tags:         []string{"perception", "object-detection", "indoor"},
			WhatsInside:  []string{"Labeled RGB images", "Bounding boxes", "Object classes", "Depth maps"},
			SizeGB:       8.3,
			SamplesCount: 5000,
			Duration:     1800,
			Source:       "uploaded",
			OwnerType:    "user",
			OwnerID:      "user-003",
			OwnerName:    "DataScience Team",
			Visibility:   "public",
		},
	}

	for _, d := range datasets {
		if err := db.Create(d).Error; err != nil {
			return fmt.Errorf("failed to create dataset: %w", err)
		}
	}

	// Create sample simulators
	simulators := []*simulator.Simulator{
		{
			Name:        "Gazebo Classic",
			Description: "Gazebo Classic simulation environment for robotics",
			Type:        "gazebo",
			Version:     "11.12.0",
			Config:      `{"physics_engine": "ODE", "render_mode": "headless", "real_time_factor": 1.0}`,
			Tags:        []string{"gazebo", "ros", "simulation"},
		},
		{
			Name:        "CARLA Simulator",
			Description: "Open-source simulator for autonomous driving research",
			Type:        "carla",
			Version:     "0.9.15",
			Config:      `{"render_quality": "Epic", "weather": "ClearNoon", "fixed_delta_seconds": 0.05}`,
			Tags:        []string{"carla", "autonomous-driving", "urban"},
		},
		{
			Name:        "Unity Robotics Hub",
			Description: "Unity-based robotics simulation platform",
			Type:        "unity",
			Version:     "2023.1.0",
			Config:      `{"graphics_api": "Vulkan", "physics_timestep": 0.02, "ros_bridge": true}`,
			Tags:        []string{"unity", "robotics", "simulation"},
		},
	}

	for _, sim := range simulators {
		if err := db.Create(sim).Error; err != nil {
			return fmt.Errorf("failed to create simulator: %w", err)
		}
	}

	// Update package counts in repositories
	for _, repo := range repos {
		var count int64
		db.Model(&pkg.Package{}).Where("repo_id = ?", repo.ID).Count(&count)
		db.Model(repo).Update("package_count", count)
	}

	log.Printf("Seed data loaded successfully: %d repos, %d packages, %d scenarios, %d datasets, %d simulators",
		len(repos), len(packages), len(scenarios), len(datasets), len(simulators))

	return nil
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
