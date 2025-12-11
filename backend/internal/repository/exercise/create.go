package exercise

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercise"
)

func (r *ExerciseRepository) Create(ctx context.Context, data *model.ExerciseListPayload) (*model.ExerciseList, error) {
	// r.conn.QueryRow("INSERT INTO ")
	return &model.ExerciseList{}, nil
}
