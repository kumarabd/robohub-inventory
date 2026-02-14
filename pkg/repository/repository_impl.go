package repository

import (
	"context"
	"gorm.io/gorm"
)

// gormRepository implements the RepoRepository interface using GORM
type gormRepository struct {
	db *gorm.DB
}

// NewRepository creates a new GORM-based repository
func NewRepository(db *gorm.DB) RepoRepository {
	return &gormRepository{db: db}
}

func (r *gormRepository) Create(ctx context.Context, repo *Repository) error {
	return r.db.WithContext(ctx).Create(repo).Error
}

func (r *gormRepository) GetByID(ctx context.Context, id string) (*Repository, error) {
	var repo Repository
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&repo).Error
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

func (r *gormRepository) GetByName(ctx context.Context, name string) (*Repository, error) {
	var repo Repository
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&repo).Error
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

func (r *gormRepository) List(ctx context.Context, limit, offset int) ([]*Repository, error) {
	var repos []*Repository
	query := r.db.WithContext(ctx)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&repos).Error
	return repos, err
}

func (r *gormRepository) Update(ctx context.Context, repo *Repository) error {
	return r.db.WithContext(ctx).Save(repo).Error
}

func (r *gormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Repository{}).Error
}
