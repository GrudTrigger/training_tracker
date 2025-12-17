package exercise

import (
	"github.com/GrudTrigger/training_tracker/backend/internal/repository"
)

type Service struct {
	repo repository.ExerciseRepo // TODO: передалать на интерфейс
}

func NewExerciseService(repo repository.ExerciseRepo) *Service {
	return &Service{
		repo: repo,
	}
}
