package trainings

import (
	"context"
	"log/slog"

	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
)

func (s *Service) GetByID(ctx context.Context, id *t.GetByIDPayload) (*t.TrainingAll, error) {
	res, err := s.repoTrainigs.GetByID(ctx, id)
	if err != nil {
		slog.Error("training by id", err)
		return nil, err
	}
	slog.Info("training by id", "training id", res.ID)
	return res, err
}
