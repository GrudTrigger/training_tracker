package exercise

import (
	"context"
	"errors"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) FindById(ctx context.Context, uuid string) (*model.Exercises, error) {
	var exercise model.Exercises
	row := r.conn.QueryRow(ctx, "SELECT id, title, muscle_group FROM exercises WHERE id=$1", uuid)
	err := row.Scan(&exercise.ID, &exercise.Title, &exercise.MuscleGroup)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.MakeNotFound(err)
		}
		return nil, err
	}
	return &exercise, nil
}
