package dataset

import (
	"context"
	"errors"
)

var (
	ErrDatasetNotFound = errors.New("dataset not found")
	ErrInvalidDataset  = errors.New("invalid dataset data")
)

// Service handles business logic for datasets
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateDataset(ctx context.Context, dataset *Dataset) error {
	if dataset.Name == "" {
		return ErrInvalidDataset
	}
	return s.repo.Create(ctx, dataset)
}

func (s *Service) GetDataset(ctx context.Context, id uint) (*Dataset, error) {
	dataset, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrDatasetNotFound
	}
	return dataset, nil
}

func (s *Service) GetDatasetByName(ctx context.Context, name string) (*Dataset, error) {
	dataset, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return nil, ErrDatasetNotFound
	}
	return dataset, nil
}

func (s *Service) ListDatasets(ctx context.Context, limit, offset int) ([]*Dataset, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *Service) UpdateDataset(ctx context.Context, dataset *Dataset) error {
	if dataset.Name == "" {
		return ErrInvalidDataset
	}
	return s.repo.Update(ctx, dataset)
}

func (s *Service) DeleteDataset(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
