package exercise

import "github.com/jackc/pgx/v5"

type ExersiceRepository struct {
	conn *pgx.Conn
}

func NewExerciseRepository(conn *pgx.Conn) *ExersiceRepository {
	return &ExersiceRepository{
		conn: conn,
	}
}
