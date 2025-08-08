package training

import (
	"database/sql"
	"errors"

	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"github.com/GrudTrigger/trainin_tracker/internal/exercise"
	"github.com/GrudTrigger/trainin_tracker/pkg/res"
	"github.com/GrudTrigger/trainin_tracker/pkg/storage"
	"github.com/GrudTrigger/trainin_tracker/pkg/utils"
)

type IRepository interface {
	Create(input InputWithUser) (*model.Training, error)
	GetAll(input model.SearchTrainings) ([]*model.Training, error)
	GetById(id string) (*model.Training, error)
	GetMyTrainings(userId string) ([]*model.Training, error)
	DeleteById(id string) (string, error)
}

type Repository struct {
	*storage.DbPostgres
	repoExercise exercise.IExerciseRepository
}

func NewTrainingRepository(db *storage.DbPostgres, repoExercise exercise.IExerciseRepository) IRepository {
	return &Repository{db, repoExercise}
}

func (r *Repository) Create(input InputWithUser) (*model.Training, error) {
	var training model.Training
	training = model.Training{
		Exercises: []*model.Exercise{},
	}
	query := `INSERT INTO training(user_id, name, duration, date, notes) 
									VALUES($1, $2, $3, $4, $5)
									RETURNING id, user_id, name,duration, date, notes, created_at`
	err := r.QueryRow(query, input.UserId, input.Title, input.Duration, input.Date, input.Notes).
		Scan(&training.ID, &training.UserID, &training.Title, &training.Duration, &training.Date, &training.Notes, &training.CreatedAt)
	if err != nil {
		return nil, err
	}

	e, err := r.repoExercise.Create(input.Exercises, training.ID)
	if err != nil {
		return nil, err
	}
	training.Exercises = e
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
		training := utils.NewTraining()
		err = rows.Scan(&training.ID, &training.UserID, &training.Title, &training.Duration, &training.Date, &training.Notes, &training.CreatedAt)
		if err != nil {
			return nil, err
		}
		trainings = append(trainings, training)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return trainings, nil
}

func (r *Repository) GetById(id string) (*model.Training, error) {
	training := utils.NewTraining()

	query := `SELECT training.*, users.id, users.email, users.login, users.password, users.role, users.created_at, exercise.* from training JOIN users ON training.user_id = users.id JOIN exercise ON training.id = exercise.training_id WHERE training.id = $1`
	row := r.QueryRow(query, id)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return nil, errors.New("тренировка не найдена")
	}

	err := row.Scan(
		&training.ID,
		&training.UserID,
		&training.Title,
		&training.Duration,
		&training.Date,
		&training.Notes,
		&training.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return training, nil
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
		err = rows.Scan(&training.ID, &training.UserID, &training.Title, &training.Duration, &training.Date, &training.Notes, &training.CreatedAt)
		if err != nil {
			return nil, err
		}
		trainings = append(trainings, &training)
	}

	return trainings, nil
}

func (r *Repository) DeleteById(id string) (string, error) {
	query := "DELETE FROM training WHERE id = $1"
	_, err := r.Exec(query, id)
	if err != nil {
		return "", err
	}
	return res.SuccessResponse, nil
}
