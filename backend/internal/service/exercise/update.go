package exercise

import (
	"context"
	"fmt"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercise"
)

func (s *ExerciseService) Update(ctx context.Context, data *model.UpdatePayload) (*model.ExerciseList, error) {
	fmt.Println("123")
	return &model.ExerciseList{}, nil
}
