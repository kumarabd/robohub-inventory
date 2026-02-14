package repository

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Repository represents a code repository in the robotics platform
// Matches API_CONTRACT.md Repository schema
type Repository struct {
	ID          string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`                                // Format: "org/repo"
	Provider    string    `gorm:"not null" json:"provider"`                                        // "github" | "gitlab" | "bitbucket"
	URL         string    `gorm:"not null" json:"url"`                                             // Full repository URL
	Description string    `json:"description,omitempty"`
	DefaultBranch string  `gorm:"not null;default:'main'" json:"defaultBranch"`
	Visibility  string    `gorm:"not null;default:'public'" json:"visibility"`                    // "public" | "private"
	
	// Sync Information
	LastSynced  time.Time `json:"lastSynced"`
	SyncStatus  string    `gorm:"not null;default:'needs_attention'" json:"syncStatus"`          // "synced" | "syncing" | "needs_attention" | "error"
	AutoSync    bool      `gorm:"default:false" json:"autoSync"`
	
	// Latest Commit Info
	LatestCommit LatestCommit `gorm:"type:jsonb" json:"latestCommit"`
	
	// Webhook Configuration
	WebhookStatus string `gorm:"default:'inactive'" json:"webhookStatus"`                         // "active" | "inactive" | "error"
	WebhookID     string `json:"webhookId,omitempty"`
	
	// Metadata
	Tags         []string `gorm:"type:text[]" json:"tags"`
	PackageCount int      `gorm:"default:0" json:"packageCount"`
	Owner        Owner    `gorm:"type:jsonb" json:"owner"`
	
	// Timestamps
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// LatestCommit represents the latest commit information
type LatestCommit struct {
	Hash    string    `json:"hash"`    // Git commit SHA
	Message string    `json:"message"`
	Author  string    `json:"author"`
	Date    time.Time `json:"date"`
	URL     string    `json:"url"`
}

// Owner represents the repository owner
type Owner struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl,omitempty"`
}

// Scan implements sql.Scanner interface for JSONB
func (c *LatestCommit) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, c)
}

// Value implements driver.Valuer interface for JSONB
func (c LatestCommit) Value() (driver.Value, error) {
	return json.Marshal(c)
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

func (Repository) TableName() string {
	return "repositories"
}
