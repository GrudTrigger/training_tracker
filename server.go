package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/GrudTrigger/trainin_tracker/internal/approach"
	"github.com/GrudTrigger/trainin_tracker/internal/exlist"
	"github.com/redis/go-redis/v9"

	"github.com/GrudTrigger/trainin_tracker/internal/exercise"
	"github.com/go-playground/validator/v10"

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

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	userRepository := user.NewRepository(dbPostgres)
	approachRepository := approach.NewRepository(dbPostgres)
	exerciseRepository := exercise.NewRepository(dbPostgres, approachRepository)
	trainingRepository := training.NewTrainingRepository(dbPostgres, exerciseRepository, rdb)
	exerciseListRepository := exlist.NewRepository(dbPostgres)

	userService := user.NewUserService(userRepository)
	trainingService := training.NewTrainingService(trainingRepository)
	exerciseService := exercise.NewService(exerciseRepository)
	exerciseListService := exlist.NewService(exerciseListRepository)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Configs:             cfg,
		UserService:         userService,
		TrainingService:     trainingService,
		ExerciseService:     exerciseService,
		ExerciseListService: exerciseListService,
		Validator:           v,
	}}))

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
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
