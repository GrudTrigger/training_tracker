package workouts_service

import (
	"context"

	"github.com/george/training-tracker/backend/internal/core/domain"
)

func (s *Service) List(ctx context.Context) ([]domain.Workout, error) {
	return s.repository.List(ctx)
}
