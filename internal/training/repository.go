package training

import (
	"database/sql"
	"errors"
	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"github.com/GrudTrigger/trainin_tracker/pkg/storage"
)

type IRepository interface {
	Create(input InputWithUser) (*model.Training, error)
	GetAll(input model.SearchTrainings) ([]*model.Training, error)
	GetById(id string) (*model.Training, error)
	GetMyTrainings(userId string) ([]*model.Training, error)
}

type Repository struct {
	*storage.DbPostgres
}

func NewTrainingRepository(db *storage.DbPostgres) IRepository {
	return &Repository{db}
}

func (r *Repository) Create(input InputWithUser) (*model.Training, error) {
	var training model.Training
	query := `INSERT INTO training(name, user_id, duration, date, notes, type) 
									VALUES($1, $2, $3, $4, $5, $6)
									RETURNING id, user_id, name, duration, date, notes, type, created_at`
	err := r.QueryRow(query, input.Name, input.UserId, input.Duration, input.Date, input.Notes, input.Type).
		Scan(&training.ID, &training.UserID, &training.Name, &training.Duration, &training.Date, &training.Notes, &training.Type, &training.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &training, nil
}

func (r *Repository) GetAll(input model.SearchTrainings) ([]*model.Training, error) {
	var trainings []*model.Training

	query, args := QueryGetAll(input)
	rows, err := r.Query(query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var training model.Training
		err = rows.Scan(&training.ID, &training.UserID, &training.Name, &training.Duration, &training.Date, &training.Notes, &training.Type, &training.CreatedAt)
		if err != nil {
			return nil, err
		}
		trainings = append(trainings, &training)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return trainings, nil
}

func (r *Repository) GetById(id string) (*model.Training, error) {
	var training model.Training
	training.UserData = &model.User{}

	query := `SELECT training.*, users.id, users.email, users.login, users.password, users.role, users.created_at from training JOIN users ON training.user_id = users.id WHERE training.id = $1`
	row := r.QueryRow(query, id)
	if row.Err() == sql.ErrNoRows {
		return nil, errors.New("тренировка не найдена")
	}

	err := row.Scan(
		&training.ID,
		&training.UserID,
		&training.Name,
		&training.Duration,
		&training.Date,
		&training.Notes,
		&training.Type,
		&training.CreatedAt,
		&training.UserData.ID,
		&training.UserData.Email,
		&training.UserData.Login,
		&training.UserData.Password,
		&training.UserData.Role,
		&training.UserData.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &training, nil
}

func (r *Repository) GetMyTrainings(userId string) ([]*model.Training, error) {

	var trainings []*model.Training

	query := "SELECT * FROM training WHERE user_id = $1"

	rows, err := r.Query(query, userId)

	if !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("тренировки не найдены")
	}

	for rows.Next() {
		var training model.Training
		err = rows.Scan(&training.ID, &training.UserID, &training.Name, &training.Duration, &training.Date, &training.Notes, &training.Type, &training.CreatedAt)
		if err != nil {
			return nil, err
		}
		trainings = append(trainings, &training)
	}

	return trainings, nil
}
