package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DbPostgres struct {
	*sql.DB
}

func NewDbPostgres(dsn string) *DbPostgres {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к DB. %s", err)
	}

	return &DbPostgres{
		db,
	}
}
