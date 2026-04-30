package statistics_service

import (
	"context"

	"github.com/george/training-tracker/backend/internal/core/domain"
)

type Repository interface {
	GetOverview(ctx context.Context) (domain.StatisticsOverview, error)
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetOverview(ctx context.Context) (domain.StatisticsOverview, error) {
	return s.repository.GetOverview(ctx)
}
