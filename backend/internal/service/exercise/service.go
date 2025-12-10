package exercise

import (
	exerciseRepo "github.com/GrudTrigger/training_tracker/backend/internal/repository/exercise"
)

type ExerciseService struct {
	repo *exerciseRepo.ExersiceRepository //TODO: передалать на интерфейс
}

func NewExerciseService(repo *exerciseRepo.ExersiceRepository) *ExerciseService {
	return &ExerciseService{
		repo: repo,
	}
}
