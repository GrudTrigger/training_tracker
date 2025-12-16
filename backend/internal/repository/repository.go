package repository

import (
	"context"

	model "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
)

type ExerciseRepo interface {
	Create(context.Context, *model.ExercisePayload) (*model.Exercises, error)
	All(context.Context, *model.AllPayload) ([]*model.Exercises, error)
	Update(context.Context, *model.UpdatePayload) (*model.Exercises, error)
	Delete(context.Context, *model.DeletePayload) error
}
