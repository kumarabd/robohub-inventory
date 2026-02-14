package simulator

import (
	"context"
	"errors"
)

var (
	ErrSimulatorNotFound = errors.New("simulator not found")
	ErrInvalidSimulator  = errors.New("invalid simulator data")
)

// Service handles business logic for simulators
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateSimulator(ctx context.Context, simulator *Simulator) error {
	if simulator.Name == "" {
		return ErrInvalidSimulator
	}
	return s.repo.Create(ctx, simulator)
}

func (s *Service) GetSimulator(ctx context.Context, id uint) (*Simulator, error) {
	simulator, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrSimulatorNotFound
	}
	return simulator, nil
}

func (s *Service) GetSimulatorByName(ctx context.Context, name string) (*Simulator, error) {
	simulator, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return nil, ErrSimulatorNotFound
	}
	return simulator, nil
}

func (s *Service) ListSimulators(ctx context.Context, limit, offset int) ([]*Simulator, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *Service) UpdateSimulator(ctx context.Context, simulator *Simulator) error {
	if simulator.Name == "" {
		return ErrInvalidSimulator
	}
	return s.repo.Update(ctx, simulator)
}

func (s *Service) DeleteSimulator(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
