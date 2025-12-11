package exercise

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercise"
)

func (s *ExerciseService) Delete(ctx context.Context, data *model.DeletePayload) error {
	return nil
}
