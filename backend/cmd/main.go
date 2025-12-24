package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	genexercise "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
	genhttpexervises "github.com/GrudTrigger/training_tracker/backend/gen/http/exercises/server"
	genhttpstatistics "github.com/GrudTrigger/training_tracker/backend/gen/http/statistics/server"
	genhttptrainings "github.com/GrudTrigger/training_tracker/backend/gen/http/trainings/server"
	genstatistics "github.com/GrudTrigger/training_tracker/backend/gen/statistics"
	gentrainings "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
	"github.com/GrudTrigger/training_tracker/backend/internal/config"
	"github.com/GrudTrigger/training_tracker/backend/internal/migrator"
	repoExercise "github.com/GrudTrigger/training_tracker/backend/internal/repository/exercise"
	repoTrainings "github.com/GrudTrigger/training_tracker/backend/internal/repository/trainings"
	"github.com/GrudTrigger/training_tracker/backend/internal/service/exercise"
	"github.com/GrudTrigger/training_tracker/backend/internal/service/statistics"
	"github.com/GrudTrigger/training_tracker/backend/internal/service/trainings"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/cors"
	goahttp "goa.design/goa/v3/http"
)

const envPath = "../.env" // TODO: после заменить на /.env потому что при билде будет лежать в корне с .envß

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	err := config.Load(envPath)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, config.AppConfig().Postgres.URI())
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	err = conn.Ping(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database in ping: %w", err))
	}

	m := migrator.NewMigrator(stdlib.OpenDB(*conn.Config().Copy()), "../migrations")
	err = m.Up()
	if err != nil {
		panic(fmt.Errorf("failed migrations: %w", err))
	}

	//-------Exersices--------
	exerciseRepo := repoExercise.NewExerciseRepository(conn)
	exerciseSvc := exercise.NewExerciseService(exerciseRepo)
	exerciseEndpoints := genexercise.NewEndpoints(exerciseSvc)

	//-------Trainigs---------
	trainingsRepo := repoTrainings.NewRepository(conn)
	trainingsSvc := trainings.NewService(trainingsRepo, exerciseRepo)
	trainingsEndpoints := gentrainings.NewEndpoints(trainingsSvc)

	//-------Statistics---------
	statisticsSvc := statistics.NewService(trainingsRepo)
	statisticsEndpoints := genstatistics.NewEndpoints(statisticsSvc)

	mux := goahttp.NewMuxer()

	exercisesHandler := genhttpexervises.New(exerciseEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)

	genhttpexervises.Mount(mux, exercisesHandler)

	trainingsHandler := genhttptrainings.New(trainingsEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)

	genhttptrainings.Mount(mux, trainingsHandler)

	statisticsHandler := genhttpstatistics.New(statisticsEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)

	genhttpstatistics.Mount(mux, statisticsHandler)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
		},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)
	port := "8080"
	server := &http.Server{Addr: ":" + port, Handler: handler}

	slog.Info("Starting server on :8080")

	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
