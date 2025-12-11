package exercise

import (
	"github.com/GrudTrigger/training_tracker/backend/internal/repository"
)

type ExerciseService struct {
	repo repository.ExerciseRepo // TODO: передалать на интерфейс
}

func NewExerciseService(repo repository.ExerciseRepo) *ExerciseService {
	return &ExerciseService{
		repo: repo,
	}
}
