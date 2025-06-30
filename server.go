package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/GrudTrigger/trainin_tracker/configs"
	"github.com/GrudTrigger/trainin_tracker/graph"
	"github.com/GrudTrigger/trainin_tracker/internal/user"
	"github.com/GrudTrigger/trainin_tracker/pkg/middleware"
	"github.com/GrudTrigger/trainin_tracker/pkg/storage"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {

	cfg := configs.LoadConfigs()

	dbPostgres := storage.NewDbPostgres(cfg.Dsn)
	defer dbPostgres.Close()

	userRepository := user.NewRepository(dbPostgres)
	userService := user.NewUserService(userRepository)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Configs: cfg, 
		UserService: userService,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.IsAuthed(srv, cfg))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
