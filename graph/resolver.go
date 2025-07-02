package graph

import (
	"github.com/GrudTrigger/trainin_tracker/configs"
	"github.com/GrudTrigger/trainin_tracker/internal/training"
	"github.com/GrudTrigger/trainin_tracker/internal/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Configs *configs.Configs
	UserService user.Service
	TrainingService training.Service
}
