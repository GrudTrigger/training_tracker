package training

import "github.com/GrudTrigger/trainin_tracker/graph/model"

type InputWithUser struct {
	model.CreateTraining
	UserId string
}

type ChanCounter struct {
	label string
	value interface{}
	err   error
}
