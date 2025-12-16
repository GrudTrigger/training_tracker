package exercise

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
)

func (s *ExerciseService) Update(ctx context.Context, data *model.UpdatePayload) (*model.Exercises, error) {
	res, err := s.repo.Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
