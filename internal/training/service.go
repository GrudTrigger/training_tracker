package training

import (
	"context"

	"github.com/GrudTrigger/trainin_tracker/graph/model"
)

type IService interface {
	Create(ctx context.Context, input model.CreateTraining, userId string) (*model.Training, error)
	FindAll(ctx context.Context, input model.SearchTrainings) ([]*model.Training, error)
	FindById(id string) (*model.Training, error)
	//GetMy(userId string) ([]*model.Training, error)
	DeleteById(id string) (string, error)
}

type Service struct {
	repo IRepository
}

func NewTrainingService(trainingRepository IRepository) IService {
	return &Service{repo: trainingRepository}
}

func (s *Service) Create(ctx context.Context, input model.CreateTraining, userId string) (*model.Training, error) {
	inputWithUser := InputWithUser{input, userId}
	t, err := s.repo.Create(ctx, inputWithUser)
	if err != nil {
		return nil, err
	}
	return t, err
}

func (s *Service) FindAll(ctx context.Context, input model.SearchTrainings) ([]*model.Training, error) {
	t, err := s.repo.GetAll(ctx, input)
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

//func (s *Service) GetMy(userId string) ([]*model.Training, error) {
//	t, err := s.repo.GetMyTrainings(userId)
//	if err != nil {
//		return nil, err
//	}
//	return t, err
//}

func (s *Service) DeleteById(id string) (string, error) {
	res, err := s.repo.DeleteById(id)
	if err != nil {
		return "", err
	}
	return res, nil
}
