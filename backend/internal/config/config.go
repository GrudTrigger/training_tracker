package config

import (
	"github.com/GrudTrigger/training_tracker/backend/internal/config/env"
	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	Postgres PostgresConfig
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	postgres, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}
	appConfig = &config{
		Postgres: postgres,
	}
	return nil
}

func AppConfig() *config {
	return appConfig
}
