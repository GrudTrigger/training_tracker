package domain

import "time"

type StatisticsOverview struct {
	WorkoutsCount  int        `json:"workoutsCount"`
	ExercisesCount int        `json:"exercisesCount"`
	TotalSets      int        `json:"totalSets"`
	TotalVolumeKg  float64    `json:"totalVolumeKg"`
	LastWorkoutAt  *time.Time `json:"lastWorkoutAt"`
}
