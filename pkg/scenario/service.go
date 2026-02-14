package scenario

import (
	"context"
	"errors"
)

var (
	ErrScenarioNotFound = errors.New("scenario not found")
	ErrInvalidScenario  = errors.New("invalid scenario data")
)

// Service handles business logic for scenarios
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateScenario(ctx context.Context, scenario *Scenario) error {
	if scenario.Name == "" {
		return ErrInvalidScenario
	}
	return s.repo.Create(ctx, scenario)
}

func (s *Service) GetScenario(ctx context.Context, id string) (*Scenario, error) {
	scenario, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrScenarioNotFound
	}
	return scenario, nil
}

func (s *Service) GetScenarioByName(ctx context.Context, name string) (*Scenario, error) {
	scenario, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return nil, ErrScenarioNotFound
	}
	return scenario, nil
}

func (s *Service) ListScenarios(ctx context.Context, limit, offset int) ([]*Scenario, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *Service) UpdateScenario(ctx context.Context, scenario *Scenario) error {
	if scenario.Name == "" {
		return ErrInvalidScenario
	}
	return s.repo.Update(ctx, scenario)
}

func (s *Service) DeleteScenario(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
