package dataset

import (
	"context"
)

// Repository defines the interface for dataset persistence
type Repository interface {
	Create(ctx context.Context, dataset *Dataset) error
	GetByID(ctx context.Context, id uint) (*Dataset, error)
	GetByName(ctx context.Context, name string) (*Dataset, error)
	List(ctx context.Context, limit, offset int) ([]*Dataset, error)
	Update(ctx context.Context, dataset *Dataset) error
	Delete(ctx context.Context, id uint) error
}
