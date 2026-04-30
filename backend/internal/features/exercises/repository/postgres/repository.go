package exercises_postgres_repository

import (
	"context"

	"github.com/george/training-tracker/backend/internal/core/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

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

func (r *Repository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(
		ctx,
		`select exists(select 1 from exercises where lower(name) = lower($1))`,
		name,
	).Scan(&exists)

	return exists, err
}

func (r *Repository) ExistsByID(ctx context.Context, id string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(
		ctx,
		`select exists(select 1 from exercises where id = $1)`,
		id,
	).Scan(&exists)

	return exists, err
}
