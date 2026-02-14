package repository

import (
	"context"
)

// Repository defines the interface for repository persistence
type RepoRepository interface {
	Create(ctx context.Context, repo *Repository) error
	GetByID(ctx context.Context, id uint) (*Repository, error)
	GetByName(ctx context.Context, name string) (*Repository, error)
	List(ctx context.Context, limit, offset int) ([]*Repository, error)
	Update(ctx context.Context, repo *Repository) error
	Delete(ctx context.Context, id uint) error
}
