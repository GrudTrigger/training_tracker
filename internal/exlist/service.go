package exlist

import "github.com/GrudTrigger/trainin_tracker/graph/model"

type IService interface {
	Create(input *model.CreateExerciseForList) (*model.ExerciseList, error)
	FindAll(input *model.GetExerciseList) ([]*model.ExerciseList, error)
	Update(input *model.UpdateExerciseForList) (*model.ExerciseList, error)
	DeleteById(id string) (string, error)
	Statistics() (*model.ExerciseListStatistic, error)
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) IService {
	return &Service{repo}
}

func (s *Service) Create(input *model.CreateExerciseForList) (*model.ExerciseList, error) {
	e, err := s.repo.Create(input)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (s *Service) FindAll(input *model.GetExerciseList) ([]*model.ExerciseList, error) {
	exList, err := s.repo.FindAll(input)
	if err != nil {
		return nil, err
	}
	return exList, nil
}

func (s *Service) Update(input *model.UpdateExerciseForList) (*model.ExerciseList, error) {
	ex, err := s.repo.Update(input)
	if err != nil {
		return nil, err
	}
	return ex, nil
}

func (s *Service) DeleteById(id string) (string, error) {
	r, err := s.repo.DeleteById(id)
	if err != nil {
		return "", err
	}
	return r, nil
}

func (s *Service) Statistics() (*model.ExerciseListStatistic, error) {
	r, err := s.repo.Statistics()
	if err != nil {
		return nil, err
	}
	return r, nil
}
