package trainings

import (
	"context"
	"fmt"
	"time"

	t "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) Create(ctx context.Context, data *t.CreateTrainingPayload) (*t.Training, error) {
	// Старт транзакции
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("starting transaction: %w", err)
	}
	// Откат транзакции
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	// Сначала создаем сущность тренировки, чтобы потом пользоваться ее UUID
	var training struct {
		ID        string
		Title     string
		Date      time.Time
		Duration  int
		CreatedAt *time.Time
	}
	err = tx.QueryRow(ctx, "INSERT INTO trainings(title, date, duration) VALUES($1, $2, $3) RETURNING id, title, date, duration, created_at", data.Title, data.Date, data.Duration).
		Scan(&training.ID, &training.Title, &training.Date, &training.Duration, &training.CreatedAt)
	if err != nil {
		return nil, err
	}

	for _, e := range data.Exercises {
		var trainingExercisesUUID string
		err = tx.QueryRow(ctx, "INSERT INTO training_exercises(training_id, exercise_id) VALUES($1, $2) RETURNING id", training.ID, e.ExerciseID).Scan(&trainingExercisesUUID)
		if err != nil {
			return nil, err
		}
		for _, s := range e.Sets {
			_, err = tx.Exec(ctx, "INSERT INTO exercise_sets(training_exercise_id, reps, weight) VALUES($1, $2, $3)", trainingExercisesUUID, s.Reps, s.Weight)
			if err != nil {
				return nil, err
			}
		}
	}
	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commiting transaction: %w", err)
	}

	c := training.CreatedAt.Format(time.DateTime)
	return &t.Training{
		ID:        training.ID,
		Title:     training.Title,
		Date:      training.Date.Format(time.DateOnly),
		Duration:  training.Duration,
		CreatedAt: &c,
	}, nil
}
