package exercise

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercise"
)

func (s *ExerciseService) All(ctx context.Context, data *model.AllPayload) ([]*model.ExerciseList, error) {
	var r []*model.ExerciseList
	return r, nil
}
