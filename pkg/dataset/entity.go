package dataset

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Dataset represents a dataset in the robotics platform
// Matches API_CONTRACT.md Dataset schema
type Dataset struct {
	ID                  string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name                string    `gorm:"uniqueIndex;not null" json:"name"`
	Slug                string    `gorm:"index" json:"slug,omitempty"`
	Description         string    `json:"description"`
	DetailedDescription string    `gorm:"type:text" json:"detailedDescription,omitempty"` // Markdown
	
	// Classification
	Type     string `gorm:"not null" json:"type"`     // "autonomous-driving" | "robotics" | "indoor-mapping" | "synthetic"
	Modality string `gorm:"not null" json:"modality"` // "camera" | "lidar" | "radar" | "imu" | "gps" | "multimodal"
	Format   string `gorm:"not null" json:"format"`   // "rosbag2" | "bag" | "parquet" | "custom"
	License  string `gorm:"not null" json:"license"`  // "MIT" | "Apache-2.0" | "CC-BY" | "CC-BY-NC" | "proprietary"
	
	// Content
	Tags       []string `gorm:"type:text[]" json:"tags"`
	WhatsInside []string `gorm:"type:text[]" json:"whatsInside"` // Bullet points
	UsageNotes string   `gorm:"type:text" json:"usageNotes,omitempty"`
	
	// Data Information
	SizeGB         float64 `json:"sizeGB"`
	SamplesCount   int     `json:"samplesCount"`
	SequencesCount int     `json:"sequencesCount,omitempty"`
	Duration       int     `json:"duration,omitempty"` // Seconds
	
	// Compatibility
	SupportedScenarios []string `gorm:"type:text[]" json:"supportedScenarios,omitempty"` // Scenario IDs
	RoboticsPlatforms  []string `gorm:"type:text[]" json:"roboticsPlatforms,omitempty"`
	
	// Metadata
	Source     string `gorm:"not null" json:"source"`     // "uploaded" | "external_link" | "partner"
	OwnerType  string `gorm:"not null" json:"ownerType"`  // "user" | "organization"
	OwnerID    string `gorm:"not null;index" json:"ownerId"`
	OwnerName  string `json:"ownerName"`
	Visibility string `gorm:"not null;default:'public'" json:"visibility"` // "public" | "private"
	
	// Preview
	PreviewAssets *PreviewAssets `gorm:"type:jsonb" json:"previewAssets,omitempty"`
	
	// Schema Information
	Schema *DatasetSchema `gorm:"type:jsonb" json:"schema,omitempty"`
	
	// Statistics
	DownloadCount int     `gorm:"default:0" json:"downloadCount"`
	UsedInRuns    int     `gorm:"default:0" json:"usedInRuns"`
	AvgRating     float64 `json:"avgRating,omitempty"`
	RatingCount   int     `json:"ratingCount,omitempty"`
	
	// Timestamps
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// PreviewAssets represents preview assets for the dataset
type PreviewAssets struct {
	ThumbnailURL string   `json:"thumbnailUrl,omitempty"`
	SampleFrames []string `json:"sampleFrames"` // URLs to preview images
	VideoPreview string   `json:"videoPreview,omitempty"`
}

// DatasetSchema represents the schema information for the dataset
type DatasetSchema struct {
	Topics     []Topic    `json:"topics"`
	DataSplits []DataSplit `json:"dataSplits,omitempty"`
}

// Topic represents a ROS topic in the dataset
type Topic struct {
	Name        string `json:"name"`
	MessageType string `json:"messageType"`
	Frequency   string `json:"frequency"`
	Description string `json:"description"`
}

// DataSplit represents a data split in the dataset
type DataSplit struct {
	Name        string  `json:"name"`
	Percentage  float64 `json:"percentage"`
	Description string  `json:"description"`
}

// Scan implements sql.Scanner interface for JSONB
func (p *PreviewAssets) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, p)
}

// Value implements driver.Valuer interface for JSONB
func (p PreviewAssets) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan implements sql.Scanner interface for JSONB
func (s *DatasetSchema) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, s)
}

// Value implements driver.Valuer interface for JSONB
func (s DatasetSchema) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (Dataset) TableName() string {
	return "datasets"
}
