package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	genexercise "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
	genhttpexervises "github.com/GrudTrigger/training_tracker/backend/gen/http/exercises/server"
	genhttptrainings "github.com/GrudTrigger/training_tracker/backend/gen/http/trainings/server"
	gentrainings "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
	"github.com/GrudTrigger/training_tracker/backend/internal/config"
	"github.com/GrudTrigger/training_tracker/backend/internal/migrator"
	repoExercise "github.com/GrudTrigger/training_tracker/backend/internal/repository/exercise"
	repoTrainings "github.com/GrudTrigger/training_tracker/backend/internal/repository/trainings"
	"github.com/GrudTrigger/training_tracker/backend/internal/service/exercise"
	"github.com/GrudTrigger/training_tracker/backend/internal/service/trainings"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/cors"
	goahttp "goa.design/goa/v3/http"
)

const envPath = "../.env" // TODO: после заменить на /.env потому что при билде будет лежать в корне с .envß

func main() {
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

	mux := goahttp.NewMuxer()

	exercisesHandler := genhttpexervises.New(exerciseEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)

	genhttpexervises.Mount(mux, exercisesHandler)

	trainingsHandler := genhttptrainings.New(trainingsEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)

	genhttptrainings.Mount(mux, trainingsHandler)

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

	for _, mount := range exercisesHandler.Mounts {
		log.Printf("%q mounted on %s %s", mount.Method, mount.Verb, mount.Pattern)
	}

	for _, mount := range trainingsHandler.Mounts {
		log.Printf("%q mounted on %s %s", mount.Method, mount.Verb, mount.Pattern)
	}

	log.Printf("Starting concerts service on :%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
