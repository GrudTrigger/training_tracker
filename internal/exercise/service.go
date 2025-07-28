package exercise

import "github.com/GrudTrigger/trainin_tracker/graph/model"

type IService interface {
	Create(input *model.CreateExercise) (*model.Exercise, error)
	GetAll(input *model.SearchExercise) ([]*model.Exercise, error)
	GetById(id string) (*model.Exercise, error)
	DeleteById(id string) (string, error)
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) IService {
	return &Service{repo}
}

func (s *Service) Create(input *model.CreateExercise) (*model.Exercise, error) {
	ex, err := s.repo.Create(input)
	if err != nil {
		return nil, err
	}
	return ex, nil
}

func (s *Service) GetAll(input *model.SearchExercise) ([]*model.Exercise, error) {
	exs, err := s.repo.GetAll(input)
	if err != nil {
		return nil, err
	}
	return exs, nil
}

func (s *Service) GetById(id string) (*model.Exercise, error){
	exs, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return exs, nil
}
func (s *Service) DeleteById(id string) (string, error) {
	res, err := s.repo.DeleteById(id)
	if err != nil {
		return "", err
	}
	return res, nil
}