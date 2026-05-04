package statistics_service

import (
	statistics_postgres_repository "github.com/george/training-tracker/backend/internal/features/statistics/repository/postgres"
)

type Service struct {
	repository statistics_postgres_repository.IRepository
}

func New(repository statistics_postgres_repository.IRepository) *Service {
	return &Service{repository: repository}
}
