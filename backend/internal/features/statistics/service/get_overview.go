package statistics_service

import (
	"context"

	"github.com/george/training-tracker/backend/internal/core/domain"
)

func (s *Service) GetOverview(ctx context.Context) (domain.StatisticsOverview, error) {
	return s.repository.GetOverview(ctx)
}
