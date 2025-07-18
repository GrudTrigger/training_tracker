package training

import "github.com/GrudTrigger/trainin_tracker/graph/model"

type IService interface {
	Create(input model.AddTraining, userId string) (*model.Training, error)
	FindAll(input model.SearchTrainings) ([]*model.Training, error)
	FindById(id string) (*model.Training, error)
}

type Service struct {
	repo Repository
}

func NewTrainingService(trainingRepository Repository) IService {
	return &Service{repo: trainingRepository}
}

func (s *Service) Create(input model.AddTraining, userId string) (*model.Training, error) {
	// нужно будет добавить валидацию на type
	inputWithUser := InputWithUser{input, userId}
	t, err := s.repo.Create(inputWithUser)
	if err != nil {
		return nil, err
	}
	return t, err
}

func (s *Service) FindAll(input model.SearchTrainings) ([]*model.Training, error) {
	t, err := s.repo.GetAll(input)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) FindById(id string) (*model.Training, error) {
	t, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return t, nil
}
