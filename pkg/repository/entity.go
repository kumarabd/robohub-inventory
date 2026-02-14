package repository

import (
	"time"
)

// Repository represents a code repository in the robotics platform
type Repository struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	URL         string    `gorm:"not null" json:"url"`
	Type        string    `gorm:"not null" json:"type"` // e.g., "git", "svn"
	Description string    `json:"description"`
	Language    string    `json:"language"`
	Tags        []string  `gorm:"type:text[]" json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Repository) TableName() string {
	return "repositories"
}
