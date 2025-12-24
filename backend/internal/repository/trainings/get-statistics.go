package trainings

import (
	"context"

	"github.com/GrudTrigger/training_tracker/backend/gen/statistics"
)

// GetStatistics TODO:После внедрения TG InitData сделать статистику под каждого пользователя
func (r *Repository) GetStatistics(ctx context.Context) (*statistics.TrainingsStatistics, error) {
	var res statistics.TrainingsStatistics

	trainingCount := r.conn.QueryRow(ctx, "SELECT COUNT(id) FROM trainings")
	err := trainingCount.Scan(&res.TrainingsCount)
	if err != nil {
		return nil, err
	}

	setsCount := r.conn.QueryRow(ctx, "SELECT COUNT(DISTINCT id) FROM exercise_sets")
	err = setsCount.Scan(&res.SetsCount)

	var repsCount int
	rows, err := r.conn.Query(ctx, "SELECT reps FROM exercise_sets")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var reps int
		err = rows.Scan(&reps)
		if err != nil {
			return nil, err
		}
		repsCount += reps
	}

	var averageDuration int
	var count int
	rows, err = r.conn.Query(ctx, "SELECT duration FROM trainings")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var d int
		err = rows.Scan(&d)
		if err != nil {
			return nil, err
		}
		averageDuration += d
		count++
	}

	res.RepsCount = repsCount
	res.AverageDuration = averageDuration / count

	return &res, nil
}
