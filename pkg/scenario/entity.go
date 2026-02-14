package scenario

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Scenario represents a test scenario in the robotics platform
// Matches API_CONTRACT.md Scenario schema
type Scenario struct {
	ID                  string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name                string    `gorm:"uniqueIndex;not null" json:"name"`
	Slug                string    `gorm:"index" json:"slug,omitempty"`                    // URL-friendly identifier
	Description         string    `json:"description"`
	DetailedDescription string    `gorm:"type:text" json:"detailedDescription,omitempty"` // Markdown
	
	// Classification
	Category     string `gorm:"not null" json:"category"`     // "navigation" | "perception" | "localization" | "planning"
	Difficulty   string `gorm:"not null" json:"difficulty"`   // "easy" | "medium" | "hard"
	MaintainedBy string `gorm:"not null" json:"maintainedBy"` // "RoboHub" | "Community" | "Partner"
	Verified     bool   `gorm:"default:false" json:"verified"`
	
	// Content
	WhatItTests      []string `gorm:"type:text[]" json:"whatItTests"`       // Bullet points
	WhyItMatters     string   `json:"whyItMatters"`
	RealWorldAnalogs []string `gorm:"type:text[]" json:"realWorldAnalogs"`
	Domain           string   `json:"domain"` // "indoor" | "outdoor" | "urban" | "warehouse" | "mixed"
	
	// Compatibility
	SupportedSimulators []string `gorm:"type:text[]" json:"supportedSimulators"` // e.g., ["Gazebo", "CARLA", "AirSim"]
	
	// Related Data
	RecommendedDatasets []string        `gorm:"type:text[]" json:"recommendedDatasets"` // Dataset IDs
	RequiredInputs      RequiredInputs  `gorm:"type:jsonb" json:"requiredInputs"`
	
	// Metrics
	SuccessCriteria SuccessCriteria `gorm:"type:jsonb" json:"successCriteria"`
	PassDefinition  string          `gorm:"type:text" json:"passDefinition"`
	
	// Statistics
	WeeklyRunCount      int     `gorm:"default:0" json:"weeklyRunCount"`
	MonthlyRunCount     int     `gorm:"default:0" json:"monthlyRunCount"`
	UsedByPackagesCount int     `gorm:"default:0" json:"usedByPackagesCount"`
	UsedByStacksCount   int     `gorm:"default:0" json:"usedByStacksCount"`
	AveragePassRate     float64 `gorm:"default:0" json:"averagePassRate"` // 0-100 percentage
	
	// Metadata
	Tags    []string `gorm:"type:text[]" json:"tags"`
	Owner   Owner    `gorm:"type:jsonb" json:"owner"`
	Version string   `json:"version"`
	
	// Timestamps
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// RequiredInput represents a required input for the scenario
type RequiredInput struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// RequiredInputs is a custom type for JSONB array
type RequiredInputs []RequiredInput

// Scan implements sql.Scanner interface for JSONB array
func (r *RequiredInputs) Scan(value interface{}) error {
	if value == nil {
		*r = RequiredInputs{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, r)
}

// Value implements driver.Valuer interface for JSONB array
func (r RequiredInputs) Value() (driver.Value, error) {
	if len(r) == 0 {
		return nil, nil
	}
	return json.Marshal(r)
}

// SuccessCriterion represents a success criterion for the scenario
type SuccessCriterion struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Threshold   string `json:"threshold"`
	Unit        string `json:"unit"`
}

// SuccessCriteria is a custom type for JSONB array
type SuccessCriteria []SuccessCriterion

// Scan implements sql.Scanner interface for JSONB array
func (s *SuccessCriteria) Scan(value interface{}) error {
	if value == nil {
		*s = SuccessCriteria{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, s)
}

// Value implements driver.Valuer interface for JSONB array
func (s SuccessCriteria) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}
	return json.Marshal(s)
}

// Owner represents the scenario owner
type Owner struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl,omitempty"`
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

func (Scenario) TableName() string {
	return "scenarios"
}
