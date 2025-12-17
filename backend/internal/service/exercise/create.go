package exercise

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
)

func (s *Service) Create(ctx context.Context, data *model.ExercisePayload) (*model.Exercises, error) {
	e, err := s.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return e, nil
}
