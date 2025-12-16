package exercise

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
)

func (r *ExerciseRepository) Delete(ctx context.Context, data *model.DeletePayload) error {
	_, err := r.conn.Exec(ctx, "DELETE FROM exercises WHERE id=$1", data.ExerciseID)
	if err != nil {
		return err
	}
	return nil
}
