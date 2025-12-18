package trainings

import (
	"context"

	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
)

func (r *Repository) Delete(ctx context.Context, data *t.DeletePayload) error {
	_, err := r.conn.Exec(ctx, "DELETE FROM trainings WHERE id=$1", data.UUID)
	if err != nil {
		return err
	}
	return nil
}
