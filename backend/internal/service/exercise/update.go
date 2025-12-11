package exercise

import (
	"context"
	"fmt"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercise"
)

func (s *ExerciseService) Update(ctx context.Context, data *model.UpdatePayload) (*model.ExerciseList, error) {
	fmt.Println("test actions")
	return &model.ExerciseList{}, nil
}
