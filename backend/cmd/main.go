package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	genexercise "github.com/GrudTrigger/training_tracker/backend/gen/exercise"
	genhttp "github.com/GrudTrigger/training_tracker/backend/gen/http/exercise/server"
	"github.com/GrudTrigger/training_tracker/backend/internal/config"
	repoExercise "github.com/GrudTrigger/training_tracker/backend/internal/repository/exercise"
	"github.com/GrudTrigger/training_tracker/backend/internal/service/exercise"

	"github.com/jackc/pgx/v5"
	goahttp "goa.design/goa/v3/http"
)

const envPath = "../.env" //TODO: после заменить на /.env потому что при билде будет лежать в корне с .envß

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

	exerciseRepo := repoExercise.NewExerciseRepository(conn)
	exerciseSvc := exercise.NewExerciseService(exerciseRepo)
	endpoints := genexercise.NewEndpoints(exerciseSvc)

	mux := goahttp.NewMuxer()

	handler := genhttp.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)

	genhttp.Mount(mux, handler)

	port := "8080"
	server := &http.Server{Addr: ":" + port, Handler: mux}

	for _, mount := range handler.Mounts {
		log.Printf("%q mounted on %s %s", mount.Method, mount.Verb, mount.Pattern)
	}

	log.Printf("Starting concerts service on :%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
