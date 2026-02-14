package pkg

import (
	"time"
)

// Package represents a software package in the robotics platform
type Package struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	Version     string    `gorm:"not null" json:"version"`
	Description string    `json:"description"`
	Type        string    `gorm:"not null" json:"type"` // e.g., "ros", "python", "cpp"
	Repository  string    `json:"repository"`
	Tags        []string  `gorm:"type:text[]" json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Package) TableName() string {
	return "packages"
}
