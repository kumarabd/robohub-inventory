package pkg

import (
	"context"
	"gorm.io/gorm"
)

// gormRepository implements the Repository interface using GORM
type gormRepository struct {
	db *gorm.DB
}

// NewRepository creates a new GORM-based package repository
func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) Create(ctx context.Context, pkg *Package) error {
	return r.db.WithContext(ctx).Create(pkg).Error
}

func (r *gormRepository) GetByID(ctx context.Context, id string) (*Package, error) {
	var pkg Package
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&pkg).Error
	if err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (r *gormRepository) GetByName(ctx context.Context, name string) (*Package, error) {
	var pkg Package
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&pkg).Error
	if err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (r *gormRepository) List(ctx context.Context, limit, offset int) ([]*Package, error) {
	var packages []*Package
	query := r.db.WithContext(ctx)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&packages).Error
	return packages, err
}

func (r *gormRepository) Update(ctx context.Context, pkg *Package) error {
	return r.db.WithContext(ctx).Save(pkg).Error
}

func (r *gormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Package{}).Error
}
