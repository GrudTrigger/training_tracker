package workouts_postgres_repository

import (
	"context"
	"time"

	"github.com/george/training-tracker/backend/internal/core/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

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
