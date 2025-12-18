package trainings

import (
	"context"

	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
)

func (s *Service) Delete(context.Context, *t.DeletePayload) (err error) {
	return nil
}
