package exercise

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
)

func (s *ExerciseService) Delete(ctx context.Context, data *model.DeletePayload) error {
	err := s.repo.Delete(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
