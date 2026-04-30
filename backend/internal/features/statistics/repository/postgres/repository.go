package statistics_postgres_repository

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

func (r *Repository) GetOverview(ctx context.Context) (domain.StatisticsOverview, error) {
	var overview domain.StatisticsOverview

	err := r.pool.QueryRow(
		ctx,
		`select
			(select count(*) from workouts) as workouts_count,
			(select count(*) from exercises) as exercises_count,
			(select count(*) from workout_sets) as total_sets,
			(select coalesce(sum(weight_kg * reps), 0) from workout_sets) as total_volume_kg,
			(select max(performed_at) from workouts) as last_workout_at`,
	).Scan(
		&overview.WorkoutsCount,
		&overview.ExercisesCount,
		&overview.TotalSets,
		&overview.TotalVolumeKg,
		&overview.LastWorkoutAt,
	)

	return overview, err
}
