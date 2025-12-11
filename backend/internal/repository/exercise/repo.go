package exercise

import (
	"github.com/GrudTrigger/training_tracker/backend/internal/repository"
	"github.com/jackc/pgx/v5"
)

type ExerciseRepository struct {
	conn *pgx.Conn
}

func NewExerciseRepository(conn *pgx.Conn) repository.ExerciseRepo {
	return &ExerciseRepository{
		conn: conn,
	}
}
