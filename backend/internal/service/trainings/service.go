package trainings

import "github.com/GrudTrigger/training_tracker/backend/internal/repository"

type Service struct {
	repoTrainigs repository.TrainingsRepo
	repoExercise repository.ExerciseRepo
}

func NewService(repoTrainigs repository.TrainingsRepo, repoExercise repository.ExerciseRepo) *Service {
	return &Service{repoTrainigs: repoTrainigs, repoExercise: repoExercise}
}