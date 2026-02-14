package pkg

import (
	"context"
	"errors"
)

var (
	ErrPackageNotFound = errors.New("package not found")
	ErrInvalidPackage  = errors.New("invalid package data")
)

// Service handles business logic for packages
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePackage(ctx context.Context, pkg *Package) error {
	if pkg.Name == "" || pkg.Version == "" {
		return ErrInvalidPackage
	}
	return s.repo.Create(ctx, pkg)
}

func (s *Service) GetPackage(ctx context.Context, id uint) (*Package, error) {
	pkg, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrPackageNotFound
	}
	return pkg, nil
}

func (s *Service) GetPackageByName(ctx context.Context, name string) (*Package, error) {
	pkg, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return nil, ErrPackageNotFound
	}
	return pkg, nil
}

func (s *Service) ListPackages(ctx context.Context, limit, offset int) ([]*Package, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *Service) UpdatePackage(ctx context.Context, pkg *Package) error {
	if pkg.Name == "" || pkg.Version == "" {
		return ErrInvalidPackage
	}
	return s.repo.Update(ctx, pkg)
}

func (s *Service) DeletePackage(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
