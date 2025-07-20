package user

import (
	"errors"

	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"github.com/GrudTrigger/trainin_tracker/pkg/utils"
)

type IService interface {
	Create(data *model.RegisterInput) (*model.User, error)
	Login(data *model.LoginInput) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type Service struct {
	repo IRepository
}

func NewUserService(userRepo IRepository) IService {
	return &Service{userRepo}
}

func (s *Service) Create(data *model.RegisterInput) (*model.User, error) {
	existedUser, err := s.repo.GetByEmail(data.Email)
	if err != nil {
		return nil, err
	}

	if existedUser != nil {
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

func (s *Service) Login(data *model.LoginInput) (*model.User, error) {
	existedUser, err := s.repo.GetByEmail(data.Email)
	if err != nil {
		return nil, err
	}
	err = utils.ComparePassword(data.Password, existedUser.Password)
	if err != nil {
		return nil, errors.New("неверный пароль")
	}

	//Сделать генерацию jwt
	return existedUser, nil
}

func (s *Service) GetByEmail(email string) (*model.User, error) {
	u, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("пользователь не найден")
	}
	return u, nil
}
