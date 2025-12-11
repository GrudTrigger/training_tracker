package exercise

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercise"
)

func (s *ExerciseService) Create(ctx context.Context, data *model.ExerciseListPayload) (*model.ExerciseList, error) {
	return &model.ExerciseList{
		ID:          "asdasd",
		Title:       "Жим лежа",
		MuscleGroup: 1,
	}, nil
}
