package trainings

import (
	"context"

	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
)

func(s *Service) Create(ctx context.Context, data *t.CreateTrainingPayload) (*t.Training, error) {
	// Проверяем валидность переданных упражнений
	for _, e := range data.Exercises {
		_, err := s.repoExercise.FindById(ctx, e.ExerciseID)
		if err != nil {
			return nil, t.MakeBadRequest(err)
		}
	}
	// Создание тренировки
	training, err := s.repoTrainigs.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return training, nil
}