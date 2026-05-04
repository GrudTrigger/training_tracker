package exercises_postgres_repository

import (
	"context"

	"github.com/george/training-tracker/backend/internal/core/domain"
)

func (r *Repository) List(ctx context.Context) ([]domain.Exercise, error) {
	rows, err := r.pool.Query(
		ctx,
		`select id, name, muscle_group, created_at from exercises order by created_at desc`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	exercises := make([]domain.Exercise, 0)
	for rows.Next() {
		var exercise domain.Exercise
		if err := rows.Scan(&exercise.ID, &exercise.Name, &exercise.MuscleGroup, &exercise.CreatedAt); err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}

	return exercises, rows.Err()
}
