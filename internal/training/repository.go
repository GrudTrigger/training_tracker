package training

import (
	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"github.com/GrudTrigger/trainin_tracker/pkg/storage"
)

type Repository interface {
	Create(input InputWithUser) (*model.Training, error)
	GetAll(input *model.SearchTrainings) ([]model.Training, error)
}

type TrainingRepository struct {
	*storage.DbPostgres
}

func NewTrainingRepository(db *storage.DbPostgres) Repository {
	return &TrainingRepository{db}
}

func (r *TrainingRepository) Create(input InputWithUser) (*model.Training, error) {
	var training model.Training
	query := `INSERT INTO training(name, user_id, duration, date, notes, type) 
									VALUES($1, $2, $3, $4, $5, $6)
									RETURNING id, user_id, name, duration, date, notes, type, created_at`
	err := r.QueryRow(query, input.Name, input.User_id, input.Duration, input.Date, input.Notes, input.Type).
		Scan(&training.ID, &training.UserID, &training.Name, &training.Duration, &training.Date, &training.Notes, &training.Type, &training.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &training, nil
}

func (r *TrainingRepository) GetAll(input *model.SearchTrainings) ([]model.Training, error) {
	var trainings []model.Training

	var searchName string
	if input.Name != nil {
		searchName = "%" + *input.Name + "%"
	} else {
		searchName = "%%"
	}
	query := `SELECT * FROM training
									WHERE name iLIKE $1 AND type = $2
									LIMIT $3 OFFSET $4`
	rows, err := r.Query(query, searchName, input.Type, input.Limit, input.Offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var training model.Training
		err := rows.Scan(&training.ID, &training.UserID, &training.Name, &training.Duration, &training.Date, &training.Notes, &training.Type, &training.CreatedAt)
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
