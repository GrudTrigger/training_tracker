package workouts_service

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/george/training-tracker/backend/internal/core/domain"
	core_errors "github.com/george/training-tracker/backend/internal/core/errors"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, workout domain.Workout) (domain.Workout, error)
	List(ctx context.Context) ([]domain.Workout, error)
}

type ExercisesRepository interface {
	ExistsByID(ctx context.Context, id string) (bool, error)
}

type Service struct {
	repository          Repository
	exercisesRepository ExercisesRepository
}

type CreateSetInput struct {
	ExerciseID string
	Reps       int
	WeightKg   float64
	SetOrder   int
}

func New(repository Repository, exercisesRepository ExercisesRepository) *Service {
	return &Service{
		repository:          repository,
		exercisesRepository: exercisesRepository,
	}
}

func (s *Service) Create(ctx context.Context, performedAt time.Time, note string, sets []CreateSetInput) (domain.Workout, error) {
	if len(sets) == 0 {
		return domain.Workout{}, fmt.Errorf("%w: workout must contain at least one set", core_errors.ErrValidation)
	}

	workoutSets := make([]domain.WorkoutSet, 0, len(sets))
	for _, set := range sets {
		if _, err := uuid.Parse(set.ExerciseID); err != nil {
			return domain.Workout{}, fmt.Errorf("%w: invalid exerciseId %q", core_errors.ErrValidation, set.ExerciseID)
		}
		exists, err := s.exercisesRepository.ExistsByID(ctx, set.ExerciseID)
		if err != nil {
			return domain.Workout{}, err
		}
		if !exists {
			return domain.Workout{}, fmt.Errorf("%w: exercise %q does not exist", core_errors.ErrValidation, set.ExerciseID)
		}
		if set.Reps <= 0 || set.WeightKg < 0 {
			return domain.Workout{}, fmt.Errorf("%w: reps must be > 0 and weightKg must be >= 0", core_errors.ErrValidation)
		}

		workoutSets = append(workoutSets, domain.WorkoutSet{
			ID:         uuid.NewString(),
			ExerciseID: set.ExerciseID,
			Reps:       set.Reps,
			WeightKg:   set.WeightKg,
			SetOrder:   set.SetOrder,
		})
	}

	sort.Slice(workoutSets, func(i int, j int) bool {
		return workoutSets[i].SetOrder < workoutSets[j].SetOrder
	})

	return s.repository.Create(ctx, domain.Workout{
		ID:          uuid.NewString(),
		PerformedAt: performedAt.UTC(),
		Note:        note,
		Sets:        workoutSets,
	})
}

func (s *Service) List(ctx context.Context) ([]domain.Workout, error) {
	return s.repository.List(ctx)
}
