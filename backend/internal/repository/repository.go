package repository

import (
	"context"

	e "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
)

type ExerciseRepo interface {
	Create(context.Context, *e.ExercisePayload) (*e.Exercises, error)
	All(context.Context, *e.AllPayload) ([]*e.Exercises, error)
	Update(context.Context, *e.UpdatePayload) (*e.Exercises, error)
	Delete(context.Context, *e.DeletePayload) error
	FindById(context.Context, string) (*e.Exercises, error)
}

type TrainingsRepo interface {
	Create(context.Context, *t.CreateTrainingPayload) (*t.Training, error)
	All(context.Context, *t.AllPayload) (res []*t.TrainingAll, err error)
	Delete(context.Context, *t.DeletePayload) (err error)
	GetByID(context.Context, *t.GetByIDPayload) (res *t.TrainingAll, err error)
}
