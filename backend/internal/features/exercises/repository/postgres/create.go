package exercises_postgres_repository

import (
	"context"

	"github.com/george/training-tracker/backend/internal/core/domain"
)

func (r *Repository) Create(ctx context.Context, exercise domain.Exercise) (domain.Exercise, error) {
	_, err := r.pool.Exec(
		ctx,
		`insert into exercises (id, name, muscle_group, created_at) values ($1, $2, $3, $4)`,
		exercise.ID,
		exercise.Name,
		exercise.MuscleGroup,
		exercise.CreatedAt,
	)
	if err != nil {
		return domain.Exercise{}, err
	}

	return exercise, nil
}
