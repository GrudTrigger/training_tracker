package exercises_service

import (
	"context"

	"github.com/george/training-tracker/backend/internal/core/domain"
)

func (s *Service) List(ctx context.Context) ([]domain.Exercise, error) {
	return s.repository.List(ctx)
}
