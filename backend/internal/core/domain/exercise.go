package domain

import "time"

type Exercise struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	MuscleGroup string    `json:"muscleGroup"`
	CreatedAt   time.Time `json:"createdAt"`
}
