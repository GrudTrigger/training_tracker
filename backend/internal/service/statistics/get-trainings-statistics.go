package statistics

import (
	"context"
	"log/slog"

	"github.com/GrudTrigger/training_tracker/backend/gen/statistics"
)

func (s *Service) GetTrainingsStatistics(ctx context.Context) (*statistics.TrainingsStatistics, error) {
	res, err := s.repoTrainings.GetStatistics(ctx)
	if err != nil {
		slog.Error("get statistics", err)
		return nil, err
	}
	slog.Info("get statistics")
	return res, nil
}
