package repository

import (
	"context"

	e "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
	"github.com/GrudTrigger/training_tracker/backend/gen/statistics"
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
	All(context.Context, *t.AllPayload) ([]*t.TrainingAll, error)
	Delete(context.Context, *t.DeletePayload) error
	GetByID(context.Context, *t.GetByIDPayload) (*t.TrainingAll, error)
	GetStatistics(context.Context) (*statistics.TrainingsStatistics, error)
}
