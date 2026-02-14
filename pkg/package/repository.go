package pkg

import (
	"context"
)

// Repository defines the interface for package persistence
type Repository interface {
	Create(ctx context.Context, pkg *Package) error
	GetByID(ctx context.Context, id uint) (*Package, error)
	GetByName(ctx context.Context, name string) (*Package, error)
	List(ctx context.Context, limit, offset int) ([]*Package, error)
	Update(ctx context.Context, pkg *Package) error
	Delete(ctx context.Context, id uint) error
}
