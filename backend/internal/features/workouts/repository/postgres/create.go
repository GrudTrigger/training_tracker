package workouts_postgres_repository

import (
	"context"

	"github.com/george/training-tracker/backend/internal/core/domain"
)

func (r *Repository) Create(ctx context.Context, workout domain.Workout) (domain.Workout, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.Workout{}, err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(
		ctx,
		`insert into workouts (id, performed_at, note) values ($1, $2, $3)`,
		workout.ID,
		workout.PerformedAt,
		workout.Note,
	)
	if err != nil {
		return domain.Workout{}, err
	}

	for _, set := range workout.Sets {
		_, err = tx.Exec(
			ctx,
			`insert into workout_sets (id, workout_id, exercise_id, reps, weight_kg, set_order) values ($1, $2, $3, $4, $5, $6)`,
			set.ID,
			workout.ID,
			set.ExerciseID,
			set.Reps,
			set.WeightKg,
			set.SetOrder,
		)
		if err != nil {
			return domain.Workout{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return domain.Workout{}, err
	}

	return workout, nil
}
