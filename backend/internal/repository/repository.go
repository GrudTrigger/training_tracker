package repository

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercise"
)

type ExerciseRepo interface {
	Create(ctx context.Context, data *model.ExerciseListPayload) (*model.ExerciseList, error)
}
