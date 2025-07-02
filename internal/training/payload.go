package training

import "github.com/GrudTrigger/trainin_tracker/graph/model"


type InputWithUser struct {
	*model.AddTraining
	User_id string
}