package trainings

import (
	"github.com/GrudTrigger/training_tracker/backend/internal/repository"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) repository.TrainingsRepo {
	return &Repository{conn: conn}
}