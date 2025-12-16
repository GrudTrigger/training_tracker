package exercise

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
)

func (r *ExerciseRepository) All(ctx context.Context, data *model.AllPayload) ([]*model.Exercises, error) {
	var res []*model.Exercises
	rows, err := r.conn.Query(ctx, "SELECT * from exercises LIMIT $1 OFFSET $2", data.Limit, data.Offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var e model.Exercises
		err = rows.Scan(&e.ID, &e.Title, &e.MuscleGroup)
		if err != nil {
			return nil, err
		}
		res = append(res, &e)
	}
	return res, nil
}
