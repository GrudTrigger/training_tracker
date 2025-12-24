package statistics

import (
	"context"

	"github.com/GrudTrigger/training_tracker/backend/gen/statistics"
)

func (s *Service) GetTrainingsStatistics(ctx context.Context) (*statistics.TrainingsStatistics, error) {
	res, err := s.repoTrainings.GetStatistics(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
