package training

import "github.com/GrudTrigger/trainin_tracker/graph/model"

type Service interface{
	Create(input *model.AddTraining, user_id string) (*model.Training, error)
}

type TrainingService struct{
	repo Repository
}

func NewTrainingService(trainingRepository Repository) Service {
	return &TrainingService{repo: trainingRepository}
}

func(s *TrainingService) Create(input *model.AddTraining, user_id string) (*model.Training, error) {
	// нужно будет добавить валидацию на type
	inputWithUser := InputWithUser{input, user_id}
	t, err := s.repo.Create(inputWithUser)
	if err != nil {
		return nil, err
	}
	return t, err
}

