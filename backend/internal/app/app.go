package app

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	genhttpexervises "github.com/GrudTrigger/training_tracker/backend/gen/http/exercises/server"
	genhttpstatistics "github.com/GrudTrigger/training_tracker/backend/gen/http/statistics/server"
	genhttptrainings "github.com/GrudTrigger/training_tracker/backend/gen/http/trainings/server"
	"github.com/rs/cors"
	goahttp "goa.design/goa/v3/http"
)

type App struct {
	Di     *DiContainer
	Mux    goahttp.ResolverMuxer
	Server *http.Server
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initDi,
		a.initLogger,
		a.initMux,
		a.initConfigServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initDi(_ context.Context) error {
	a.Di = NewDiContainer()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)
	return nil
}

func (a *App) initMux(ctx context.Context) error {
	mux := goahttp.NewMuxer()

	exercisesHandler := genhttpexervises.New(a.Di.ExerciseEndpoints(ctx), mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)

	genhttpexervises.Mount(mux, exercisesHandler)

	trainingsHandler := genhttptrainings.New(a.Di.TrainingsEndpoints(ctx), mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)

	genhttptrainings.Mount(mux, trainingsHandler)

	statisticsHandler := genhttpstatistics.New(a.Di.StatisticsEndpoints(ctx), mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)

	genhttpstatistics.Mount(mux, statisticsHandler)

	a.Mux = mux

	return nil
}

func (a *App) initConfigServer(_ context.Context) error {
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

	handler := c.Handler(a.Mux)
	port := "8080"
	server := &http.Server{Addr: ":" + port, Handler: handler}
	a.Server = server

	return nil
}

func (a *App) Run(_ context.Context) error {
	slog.Info(fmt.Sprintf("HTTP сервер запущен на порту %s\n", "8080"))
	defer a.Server.Close()
	if err := a.Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return nil
}
