package scenario

import (
	"context"
	"gorm.io/gorm"
)

// gormRepository implements the Repository interface using GORM
type gormRepository struct {
	db *gorm.DB
}

// NewRepository creates a new GORM-based scenario repository
func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) Create(ctx context.Context, scenario *Scenario) error {
	return r.db.WithContext(ctx).Create(scenario).Error
}

func (r *gormRepository) GetByID(ctx context.Context, id string) (*Scenario, error) {
	var scenario Scenario
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&scenario).Error
	if err != nil {
		return nil, err
	}
	return &scenario, nil
}

func (r *gormRepository) GetByName(ctx context.Context, name string) (*Scenario, error) {
	var scenario Scenario
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&scenario).Error
	if err != nil {
		return nil, err
	}
	return &scenario, nil
}

func (r *gormRepository) List(ctx context.Context, limit, offset int) ([]*Scenario, error) {
	var scenarios []*Scenario
	query := r.db.WithContext(ctx)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&scenarios).Error
	return scenarios, err
}

func (r *gormRepository) Update(ctx context.Context, scenario *Scenario) error {
	return r.db.WithContext(ctx).Save(scenario).Error
}

func (r *gormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Scenario{}).Error
}
