package domain

import "time"

type Workout struct {
	ID          string       `json:"id"`
	PerformedAt time.Time    `json:"performedAt"`
	Note        string       `json:"note"`
	Sets        []WorkoutSet `json:"sets"`
}

type WorkoutSet struct {
	ID         string  `json:"id"`
	ExerciseID string  `json:"exerciseId"`
	Reps       int     `json:"reps"`
	WeightKg   float64 `json:"weightKg"`
	SetOrder   int     `json:"setOrder"`
}
