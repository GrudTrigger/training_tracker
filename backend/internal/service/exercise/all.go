package exercise

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
)

func (s *ExerciseService) All(ctx context.Context, data *model.AllPayload) ([]*model.Exercises, error) {
	res, err := s.repo.All(ctx, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
