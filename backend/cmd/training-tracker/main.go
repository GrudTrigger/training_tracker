package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/george/training-tracker/backend/internal/core/config"
	core_httpserver "github.com/george/training-tracker/backend/internal/core/httpserver"
	core_logger "github.com/george/training-tracker/backend/internal/core/logger"
	core_postgres "github.com/george/training-tracker/backend/internal/core/postgres"
	core_redis "github.com/george/training-tracker/backend/internal/core/redis"
	exercises_postgres "github.com/george/training-tracker/backend/internal/features/exercises/repository/postgres"
	exercises_service "github.com/george/training-tracker/backend/internal/features/exercises/service"
	exercises_http "github.com/george/training-tracker/backend/internal/features/exercises/transport/http"
	statistics_postgres "github.com/george/training-tracker/backend/internal/features/statistics/repository/postgres"
	statistics_service "github.com/george/training-tracker/backend/internal/features/statistics/service"
	statistics_http "github.com/george/training-tracker/backend/internal/features/statistics/transport/http"
	workouts_postgres "github.com/george/training-tracker/backend/internal/features/workouts/repository/postgres"
	workouts_service "github.com/george/training-tracker/backend/internal/features/workouts/service"
	workouts_http "github.com/george/training-tracker/backend/internal/features/workouts/transport/http"
)

func main() {
	cfg, err := core_config.New()
	if err != nil {
		fmt.Println("load config:", err)
		os.Exit(1)
	}

	logger := core_logger.New(cfg.AppEnv)
	time.Local = cfg.TimeZone

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	postgresPool, err := core_postgres.NewPool(ctx, cfg.PostgresDSN)
	if err != nil {
		logger.Error("connect postgres", "error", err)
		os.Exit(1)
	}
	defer postgresPool.Close()

	redisClient, err := core_redis.NewClient(ctx, cfg.RedisAddr, cfg.RedisDB)
	if err != nil {
		logger.Error("connect redis", "error", err)
		os.Exit(1)
	}
	defer redisClient.Close()

	_ = redisClient

	logger.Debug("initializing feature", "feature", "exercises")
	exercisesRepository := exercises_postgres.NewRepository(postgresPool)
	exercisesHandler := exercises_http.NewHandler(
		exercises_service.New(exercisesRepository),
	)
	logger.Debug("initializing feature", "feature", "workouts")
	workoutsRepository := workouts_postgres.NewRepository(postgresPool)
	workoutsHandler := workouts_http.NewHandler(
		workouts_service.New(workoutsRepository, exercisesRepository),
	)

	logger.Debug("initializing feature", "feature", "statistics")
	statisticsRepository := statistics_postgres.NewRepository(postgresPool)
	statisticsHandler := statistics_http.NewHandler(
		statistics_service.New(statisticsRepository),
	)

	logger.Debug("initializing HTTP server")
	server := core_httpserver.New(cfg.HTTPHost, cfg.HTTPPort, logger)
	server.Register("GET", "/health", func(w http.ResponseWriter, _ *http.Request) {
		core_httpserver.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	api := core_httpserver.NewGroup("/api/v1")
	exercisesHandler.Register(api)
	workoutsHandler.Register(api)
	statisticsHandler.Register(api)
	server.RegisterGroup(api)

	logger.Info("starting HTTP server", "addr", server.Address())

	if err := server.Run(ctx); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
