package simulator

import "context"

// Repository defines the interface for simulator persistence
type Repository interface {
	Create(ctx context.Context, simulator *Simulator) error
	GetByID(ctx context.Context, id string) (*Simulator, error)
	GetByName(ctx context.Context, name string) (*Simulator, error)
	List(ctx context.Context, limit, offset int) ([]*Simulator, error)
	Update(ctx context.Context, simulator *Simulator) error
	Delete(ctx context.Context, id string) error
}
