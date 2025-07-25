package main

import (
	"flag"
	"fmt"
	"github.com/GrudTrigger/trainin_tracker/internal/exercise"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/GrudTrigger/trainin_tracker/configs"
	"github.com/GrudTrigger/trainin_tracker/graph"
	"github.com/GrudTrigger/trainin_tracker/internal/training"
	"github.com/GrudTrigger/trainin_tracker/internal/user"
	"github.com/GrudTrigger/trainin_tracker/pkg/middleware"
	"github.com/GrudTrigger/trainin_tracker/pkg/storage"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	env := flag.String("env", "dev", "Application environment: dev or prod")
	flag.Parse()
	fmt.Println(*env)
	cfg := configs.LoadConfigs()
	v := validator.New(validator.WithRequiredStructEnabled())
	dbPostgres := storage.NewDbPostgres(cfg.Dsn)
	defer dbPostgres.Close()

	userRepository := user.NewRepository(dbPostgres)
	trainingRepository := training.NewTrainingRepository(dbPostgres)
	exerciseRepository := exercise.NewRepository(dbPostgres)

	userService := user.NewUserService(userRepository)
	trainingService := training.NewTrainingService(trainingRepository)
	exerciseService := exercise.NewService(exerciseRepository)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Configs:         cfg,
		UserService:     userService,
		TrainingService: trainingService,
		ExerciseService: exerciseService,
		Validator:       v,
	}}))

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	}).Handler)
	router.Use(middleware.Logging)
	router.Handle("/query", middleware.IsAuthed(srv, cfg))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(middleware.LoggingExtension{})
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
