package exercise

import (
	"context"
	"errors"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) Update(ctx context.Context, data *model.UpdatePayload) (*model.Exercises, error) {
	var e model.Exercises
	row := r.conn.QueryRow(ctx, "SELECT * FROM exercises WHERE id=$1", data.ID)
	err := row.Scan(&e.ID, &e.Title, &e.MuscleGroup)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.MakeNotFound(err)
		}
		return nil, err
	}
	var res model.Exercises
	updateRow := r.conn.QueryRow(ctx, "UPDATE exercises SET title=$1, muscle_group=$2 WHERE id=$3 RETURNING id, title, muscle_group", data.Title, data.MuscleGroup, data.ID)
	err = updateRow.Scan(&res.ID, &res.Title, &res.MuscleGroup)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
