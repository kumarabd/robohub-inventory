package scenario

import "context"

// Repository defines the interface for scenario persistence
type Repository interface {
	Create(ctx context.Context, scenario *Scenario) error
	GetByID(ctx context.Context, id string) (*Scenario, error)
	GetByName(ctx context.Context, name string) (*Scenario, error)
	List(ctx context.Context, limit, offset int) ([]*Scenario, error)
	Update(ctx context.Context, scenario *Scenario) error
	Delete(ctx context.Context, id string) error
}
