package workouts_postgres_repository

import (
	"context"
	"time"

	"github.com/george/training-tracker/backend/internal/core/domain"
)

func (r *Repository) List(ctx context.Context) ([]domain.Workout, error) {
	rows, err := r.pool.Query(
		ctx,
		`select
			w.id,
			w.performed_at,
			w.note,
			ws.id,
			ws.exercise_id,
			ws.reps,
			ws.weight_kg,
			ws.set_order
		from workouts w
		left join workout_sets ws on ws.workout_id = w.id
		order by w.performed_at desc, ws.set_order asc`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workouts := make([]domain.Workout, 0)
	indexByID := make(map[string]int)

	for rows.Next() {
		var (
			workoutID   string
			performedAt time.Time
			note        string
			setID       *string
			exerciseID  *string
			reps        *int
			weightKg    *float64
			setOrder    *int
		)

		if err := rows.Scan(&workoutID, &performedAt, &note, &setID, &exerciseID, &reps, &weightKg, &setOrder); err != nil {
			return nil, err
		}

		index, exists := indexByID[workoutID]
		if !exists {
			workouts = append(workouts, domain.Workout{
				ID:          workoutID,
				PerformedAt: performedAt.UTC(),
				Note:        note,
				Sets:        []domain.WorkoutSet{},
			})
			index = len(workouts) - 1
			indexByID[workoutID] = index
		}

		if setID != nil && exerciseID != nil && reps != nil && weightKg != nil && setOrder != nil {
			workouts[index].Sets = append(workouts[index].Sets, domain.WorkoutSet{
				ID:         *setID,
				ExerciseID: *exerciseID,
				Reps:       *reps,
				WeightKg:   *weightKg,
				SetOrder:   *setOrder,
			})
		}
	}

	return workouts, rows.Err()
}
