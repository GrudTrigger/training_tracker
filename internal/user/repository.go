package user

import (
	"database/sql"
	"errors"

	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"github.com/GrudTrigger/trainin_tracker/pkg/storage"
)

type Repository interface {
	Create(data *model.RegisterInput) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll()
}

type UserRepository struct {
	*storage.DbPostgres
}

func NewRepository(db *storage.DbPostgres) Repository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(data *model.RegisterInput) (*model.User, error) {
	var userModel model.User
	query := `INSERT INTO users (email, login, password, role) 
			  VALUES ($1, $2, $3, $4) 
			  RETURNING id, email, login, password, role, created_at`
	err := r.QueryRow(query, data.Email, data.Login, data.Password, data.Role).
		Scan(&userModel.ID, &userModel.Email, &userModel.Login, &userModel.Password, &userModel.Role, &userModel.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &userModel, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var existedUser model.User

	query := "SELECT id, email, login, password, role FROM users WHERE email = $1"
	row := r.QueryRow(query, email)
	err := row.Scan(&existedUser.ID, &existedUser.Email, &existedUser.Login, &existedUser.Password, &existedUser.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &existedUser, nil
}

func (r *UserRepository) GetAll() {}
