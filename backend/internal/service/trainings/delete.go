package trainings

import (
	"context"
	"log/slog"

	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
)

func (s *Service) Delete(ctx context.Context, data *t.DeletePayload) error {
	err := s.repoTrainigs.Delete(ctx, data)
	if err != nil {
		slog.Error("delete training", "training id", data.UUID, err)
		return err
	}
	slog.Info("delete training", "training id", data.UUID)
	return nil
}
