package dataset

import (
	"time"
)

// Dataset represents a dataset in the robotics platform
type Dataset struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	Description string    `json:"description"`
	Type        string    `gorm:"not null" json:"type"` // e.g., "sensor", "image", "lidar", "training"
	Size        int64     `json:"size"` // Size in bytes
	Location    string    `json:"location"` // Storage location/URL
	Format      string    `json:"format"` // e.g., "rosbag", "hdf5", "json"
	Tags        []string  `gorm:"type:text[]" json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Dataset) TableName() string {
	return "datasets"
}
