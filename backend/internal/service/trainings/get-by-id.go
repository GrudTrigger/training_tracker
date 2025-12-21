package trainings

import (
	"context"

	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
)

func (s *Service) GetByID(ctx context.Context, id *t.GetByIDPayload) (*t.TrainingAll, error) {
	res, err := s.repoTrainigs.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, err
}
