package exercise

import (
	"github.com/GrudTrigger/trainin_tracker/graph/model"
	"github.com/GrudTrigger/trainin_tracker/pkg/storage"
)

type IRepository interface {
	Create(input *model.CreateExercise) (*model.Exercise, error)
}

type Repository struct {
	*storage.DbPostgres
}

func NewRepository(db *storage.DbPostgres) IRepository {
	return &Repository{db}
}

func (r *Repository) Create(input *model.CreateExercise) (*model.Exercise, error) {
	var e model.Exercise

	query := `INSERT INTO exercise(title, muscle_group, approach_count, weight, training_id) VALUES ($1, $2, $3, $4, $5)
				RETURNING id, title, muscle_group, approach_count, weight, created_at`
	err := r.QueryRow(query, input.Title, input.MuscleGroup, input.ApproachCount, input.Weight, input.TrainingID).
		Scan(e.ID, e.Title, e.MuscleGroup, e.ApproachCount, e.Weight, e.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
