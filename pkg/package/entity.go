package pkg

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Package represents a software package in the robotics platform
// Matches API_CONTRACT.md Package schema
type Package struct {
	ID          string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`                     // Package name (lowercase, no spaces)
	DisplayName string    `json:"displayName"`                                          // Human-readable name
	Description string    `json:"description"`
	Documentation string  `json:"documentation,omitempty"`                              // Markdown content or URL
	
	// Repository Information
	RepoID   string `gorm:"index" json:"repoId"`
	RepoName string `json:"repoName"`                                                  // org/repo format
	Path     string `json:"path"`                                                      // Path within repo
	
	// Type Classification
	Types []string `gorm:"type:text[]" json:"types"`                                  // "planner" | "perception" | "control" | "sensors" | "simulation" | "infrastructure" | "other"
	
	// Version Information
	LatestVersion string   `json:"latestVersion"`                                      // Semantic version
	Versions      []string `gorm:"type:text[]" json:"versions"`                        // All available versions
	
	// Metadata
	Tags     []string `gorm:"type:text[]" json:"tags"`
	Keywords []string `gorm:"type:text[]" json:"keywords"`
	
	// Validation Status
	ValidationStatus ValidationStatus `gorm:"type:jsonb" json:"validationStatus"`
	
	// Relationships
	LinkedScenariosCount    int `gorm:"default:0" json:"linkedScenariosCount"`
	LinkedDatasetsCount     int `gorm:"default:0" json:"linkedDatasetsCount"`
	UsedInCollectionsCount  int `gorm:"default:0" json:"usedInCollectionsCount"`
	
	// Owner Information
	Owner Owner `gorm:"type:jsonb" json:"owner"`
	
	// Last Run
	LastRun *LastRun `gorm:"type:jsonb" json:"lastRun,omitempty"`
	
	// License & Dependencies
	License      string       `json:"license,omitempty"`
	Dependencies Dependencies `gorm:"type:jsonb" json:"dependencies,omitempty"`
	
	// Timestamps
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ValidationStatus represents package validation information
type ValidationStatus struct {
	LastValidated time.Time `json:"lastValidated"`
	Status        string    `json:"status"`   // "pass" | "fail" | "pending"
	PassRate      float64   `json:"passRate"` // 0-100 percentage
}

// Owner represents the package owner
type Owner struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl,omitempty"`
}

// LastRun represents the last run information
type LastRun struct {
	Status     string    `json:"status"`     // "pass" | "fail" | "pending"
	RunAt      time.Time `json:"runAt"`
	ScenarioID string    `json:"scenarioId"`
}

// Dependency represents a package dependency
type Dependency struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Dependencies is a custom type for JSONB array of dependencies
type Dependencies []Dependency

// Scan implements sql.Scanner interface for JSONB array
func (d *Dependencies) Scan(value interface{}) error {
	if value == nil {
		*d = Dependencies{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, d)
}

// Value implements driver.Valuer interface for JSONB array
func (d Dependencies) Value() (driver.Value, error) {
	if len(d) == 0 {
		return nil, nil
	}
	return json.Marshal(d)
}

// Scan implements sql.Scanner interface for JSONB
func (v *ValidationStatus) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, v)
}

// Value implements driver.Valuer interface for JSONB
func (v ValidationStatus) Value() (driver.Value, error) {
	return json.Marshal(v)
}

// Scan implements sql.Scanner interface for JSONB
func (o *Owner) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, o)
}

// Value implements driver.Valuer interface for JSONB
func (o Owner) Value() (driver.Value, error) {
	return json.Marshal(o)
}

// Scan implements sql.Scanner interface for JSONB
func (l *LastRun) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, l)
}

// Value implements driver.Valuer interface for JSONB
func (l LastRun) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (Package) TableName() string {
	return "packages"
}
