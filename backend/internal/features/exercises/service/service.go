package exercises_service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/george/training-tracker/backend/internal/core/domain"
	core_errors "github.com/george/training-tracker/backend/internal/core/errors"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, exercise domain.Exercise) (domain.Exercise, error)
	List(ctx context.Context) ([]domain.Exercise, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, name string, muscleGroup string) (domain.Exercise, error) {
	name = strings.TrimSpace(name)
	muscleGroup = strings.TrimSpace(muscleGroup)

	if name == "" || muscleGroup == "" {
		return domain.Exercise{}, fmt.Errorf("%w: name and muscleGroup are required", core_errors.ErrValidation)
	}

	exists, err := s.repository.ExistsByName(ctx, name)
	if err != nil {
		return domain.Exercise{}, err
	}
	if exists {
		return domain.Exercise{}, fmt.Errorf("%w: exercise with name %q already exists", core_errors.ErrConflict, name)
	}

	return s.repository.Create(ctx, domain.Exercise{
		ID:          uuid.NewString(),
		Name:        name,
		MuscleGroup: muscleGroup,
		CreatedAt:   time.Now().UTC(),
	})
}

func (s *Service) List(ctx context.Context) ([]domain.Exercise, error) {
	return s.repository.List(ctx)
}
