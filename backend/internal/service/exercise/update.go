package exercise

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercise"
)

func (s *ExerciseService) Update(ctx context.Context, data *model.UpdatePayload) (*model.ExerciseList, error) {
	return &model.ExerciseList{}, nil
}
