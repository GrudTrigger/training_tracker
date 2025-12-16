package exercise

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
)

func (r *ExerciseRepository) Create(ctx context.Context, data *model.ExercisePayload) (*model.Exercises, error) {
	var e model.Exercises
	row := r.conn.QueryRow(ctx, "INSERT INTO exercises(title, muscle_group) VALUES($1, $2) RETURNING id, title, muscle_group", data.Title, data.MuscleGroup)
	err := row.Scan(&e.ID, &e.Title, &e.MuscleGroup)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
