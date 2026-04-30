package core_config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppEnv      string
	TimeZone    *time.Location
	HTTPHost    string
	HTTPPort    string
	PostgresDSN string
	RedisAddr   string
	RedisDB     int
}

func New() (*Config, error) {
	appEnv := valueOrDefault("APP_ENV", "development")
	timeZoneName := valueOrDefault("TIME_ZONE", "UTC")
	httpHost := valueOrDefault("HTTP_HOST", "0.0.0.0")
	httpPort := valueOrDefault("HTTP_PORT", "8080")
	postgresDSN := valueOrDefault("POSTGRES_DSN", "postgres://postgres:postgres@localhost:5432/training_tracker?sslmode=disable")
	redisAddr := valueOrDefault("REDIS_ADDR", "localhost:6379")
	redisDBRaw := valueOrDefault("REDIS_DB", "0")

	timeZone, err := time.LoadLocation(timeZoneName)
	if err != nil {
		return nil, fmt.Errorf("load time zone %q: %w", timeZoneName, err)
	}

	redisDB, err := strconv.Atoi(redisDBRaw)
	if err != nil {
		return nil, fmt.Errorf("parse redis db: %w", err)
	}

	return &Config{
		AppEnv:      appEnv,
		TimeZone:    timeZone,
		HTTPHost:    httpHost,
		HTTPPort:    httpPort,
		PostgresDSN: postgresDSN,
		RedisAddr:   redisAddr,
		RedisDB:     redisDB,
	}, nil
}

func valueOrDefault(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
