package simulator

import (
	"time"
)

// Simulator represents a simulation environment in the robotics platform
type Simulator struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	Description string    `json:"description"`
	Type        string    `gorm:"not null" json:"type"` // e.g., "gazebo", "unity", "custom"
	Version     string    `json:"version"`
	Config      string    `gorm:"type:text" json:"config"` // JSON configuration
	Tags        []string  `gorm:"type:text[]" json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Simulator) TableName() string {
	return "simulators"
}
