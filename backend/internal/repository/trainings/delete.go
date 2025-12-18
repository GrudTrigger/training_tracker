package trainings

import (
	"context"

	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
)

func (r *Repository) Delete(ctx context.Context, data *t.DeletePayload) error {
	return nil
}
