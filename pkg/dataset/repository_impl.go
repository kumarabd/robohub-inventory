package dataset

import (
	"context"
	"gorm.io/gorm"
)

// gormRepository implements the Repository interface using GORM
type gormRepository struct {
	db *gorm.DB
}

// NewRepository creates a new GORM-based dataset repository
func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) Create(ctx context.Context, dataset *Dataset) error {
	return r.db.WithContext(ctx).Create(dataset).Error
}

func (r *gormRepository) GetByID(ctx context.Context, id string) (*Dataset, error) {
	var dataset Dataset
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&dataset).Error
	if err != nil {
		return nil, err
	}
	return &dataset, nil
}

func (r *gormRepository) GetByName(ctx context.Context, name string) (*Dataset, error) {
	var dataset Dataset
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&dataset).Error
	if err != nil {
		return nil, err
	}
	return &dataset, nil
}

func (r *gormRepository) List(ctx context.Context, limit, offset int) ([]*Dataset, error) {
	var datasets []*Dataset
	query := r.db.WithContext(ctx)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&datasets).Error
	return datasets, err
}

func (r *gormRepository) Update(ctx context.Context, dataset *Dataset) error {
	return r.db.WithContext(ctx).Save(dataset).Error
}

func (r *gormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Dataset{}).Error
}
