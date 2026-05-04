package workouts_service

import (
	"context"

	workouts_postgres_repository "github.com/george/training-tracker/backend/internal/features/workouts/repository/postgres"
)

type ExercisesRepository interface {
	ExistsByID(ctx context.Context, id string) (bool, error)
}

type Service struct {
	repository          workouts_postgres_repository.IRepository
	exercisesRepository ExercisesRepository
}

type CreateSetInput struct {
	ExerciseID string
	Reps       int
	WeightKg   float64
	SetOrder   int
}

func New(repository workouts_postgres_repository.IRepository, exercisesRepository ExercisesRepository) *Service {
	return &Service{
		repository:          repository,
		exercisesRepository: exercisesRepository,
	}
}
