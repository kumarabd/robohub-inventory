package scenario

import (
	"time"
)

// Scenario represents a test scenario in the robotics platform
type Scenario struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	Description string    `json:"description"`
	Type        string    `gorm:"not null" json:"type"` // e.g., "simulation", "real_world"
	Config      string    `gorm:"type:text" json:"config"` // JSON configuration
	Tags        []string  `gorm:"type:text[]" json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Scenario) TableName() string {
	return "scenarios"
}
