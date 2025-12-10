package exersice

import (
	"context"

	gen "github.com/GrudTrigger/training_tracker/backend/gen/exercise"
)

func (s *ExerciseService) Create(context.Context, *gen.ExerciseListPayload) (res *gen.ExerciseList, err error) {
	return &gen.ExerciseList{
		ID:          "asdasd",
		Title:       "Жим лежа",
		MuscleGroup: 1,
	}, nil
}
