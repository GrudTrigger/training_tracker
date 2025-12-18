package trainings

import (
	"context"

	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
)

func (s *Service) All(ctx context.Context, data *t.AllPayload) ([]*t.TrainingAll, error) {
	res, err := s.repoTrainigs.All(ctx, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
