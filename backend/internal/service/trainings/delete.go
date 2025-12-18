package trainings

import (
	"context"

	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
)

func (s *Service) Delete(ctx context.Context, data *t.DeletePayload) error {
	err := s.repoTrainigs.Delete(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
