package statistics

import "github.com/GrudTrigger/training_tracker/backend/internal/repository"

type Service struct {
	repoTrainings repository.TrainingsRepo
}

func NewService(repoTrainings repository.TrainingsRepo) *Service {
	return &Service{
		repoTrainings: repoTrainings,
	}
}
