package exercise

import "github.com/GrudTrigger/trainin_tracker/graph/model"

type IService interface {
	Create(input *model.CreateExercise) (*model.Exercise, error)
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
