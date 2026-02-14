package simulator

import (
	"context"
	"gorm.io/gorm"
)

// gormRepository implements the Repository interface using GORM
type gormRepository struct {
	db *gorm.DB
}

// NewRepository creates a new GORM-based repository
func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) Create(ctx context.Context, simulator *Simulator) error {
	return r.db.WithContext(ctx).Create(simulator).Error
}

func (r *gormRepository) GetByID(ctx context.Context, id uint) (*Simulator, error) {
	var simulator Simulator
	err := r.db.WithContext(ctx).First(&simulator, id).Error
	if err != nil {
		return nil, err
	}
	return &simulator, nil
}

func (r *gormRepository) GetByName(ctx context.Context, name string) (*Simulator, error) {
	var simulator Simulator
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&simulator).Error
	if err != nil {
		return nil, err
	}
	return &simulator, nil
}

func (r *gormRepository) List(ctx context.Context, limit, offset int) ([]*Simulator, error) {
	var simulators []*Simulator
	query := r.db.WithContext(ctx)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Find(&simulators).Error
	return simulators, err
}

func (r *gormRepository) Update(ctx context.Context, simulator *Simulator) error {
	return r.db.WithContext(ctx).Save(simulator).Error
}

func (r *gormRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Simulator{}, id).Error
}
