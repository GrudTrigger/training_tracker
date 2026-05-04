package exercises_postgres_repository

import (
	"context"

	"github.com/george/training-tracker/backend/internal/core/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IRepository interface {
	Create(ctx context.Context, exercise domain.Exercise) (domain.Exercise, error)
	List(ctx context.Context) ([]domain.Exercise, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
}

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}
