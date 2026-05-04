package exercises_service

import (
	exercises_postgres_repository "github.com/george/training-tracker/backend/internal/features/exercises/repository/postgres"
)

type Service struct {
	repository exercises_postgres_repository.IRepository
}

func New(repository exercises_postgres_repository.IRepository) *Service {
	return &Service{repository: repository}
}
