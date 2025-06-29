package user

import (
	"errors"
	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"github.com/GrudTrigger/trainin_tracker/pkg/utils"
)

type Service interface {
	Create(data *CreateRequest) (*model.User, error)
}

type UserService struct {
	repo Repository
}

func NewUserService(userRepo Repository) Service {
	return &UserService{userRepo}
}

func (s *UserService) Create(data *CreateRequest) (*model.User, error) {
	existedUser, err := s.repo.GetByEmail(data.Email)
	if err != nil {
		return nil, err
	}

	if existedUser != "" {
		return nil, errors.New("пользователь с такой почтой уже зарегистрирован")
	}

	hashedPassword, err := utils.HashedPassword(data.Password)
	if err != nil {
		return nil, err
	}

	data.Password = hashedPassword

	newUser, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
