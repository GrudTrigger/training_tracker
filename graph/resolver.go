package graph

import (
	"github.com/GrudTrigger/trainin_tracker/configs"
	"github.com/GrudTrigger/trainin_tracker/internal/exercise"
	"github.com/GrudTrigger/trainin_tracker/internal/exlist"
	"github.com/GrudTrigger/trainin_tracker/internal/training"
	"github.com/GrudTrigger/trainin_tracker/internal/user"
	"github.com/go-playground/validator/v10"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Configs             *configs.Configs
	UserService         user.IService
	TrainingService     training.IService
	ExerciseService     exercise.IService
	ExerciseListService exlist.IService
	Validator           *validator.Validate
}
