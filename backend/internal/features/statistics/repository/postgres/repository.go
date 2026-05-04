package statistics_postgres_repository

import (
	"context"

	"github.com/george/training-tracker/backend/internal/core/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IRepository interface {
	GetOverview(ctx context.Context) (domain.StatisticsOverview, error)
}

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}
