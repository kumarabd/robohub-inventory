package repository

import (
	"context"
	"errors"
)

var (
	ErrRepositoryNotFound = errors.New("repository not found")
	ErrInvalidRepository  = errors.New("invalid repository data")
)

// Service handles business logic for repositories
type Service struct {
	repo RepoRepository
}

func NewService(repo RepoRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateRepository(ctx context.Context, repo *Repository) error {
	if repo.Name == "" || repo.URL == "" {
		return ErrInvalidRepository
	}
	return s.repo.Create(ctx, repo)
}

func (s *Service) GetRepository(ctx context.Context, id string) (*Repository, error) {
	repo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrRepositoryNotFound
	}
	return repo, nil
}

func (s *Service) GetRepositoryByName(ctx context.Context, name string) (*Repository, error) {
	repo, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return nil, ErrRepositoryNotFound
	}
	return repo, nil
}

func (s *Service) ListRepositories(ctx context.Context, limit, offset int) ([]*Repository, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *Service) UpdateRepository(ctx context.Context, repo *Repository) error {
	if repo.Name == "" || repo.URL == "" {
		return ErrInvalidRepository
	}
	return s.repo.Update(ctx, repo)
}

func (s *Service) DeleteRepository(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
