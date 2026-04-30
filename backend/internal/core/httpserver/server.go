package core_httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Group struct {
	prefix string
	routes []route
}

func NewGroup(prefix string) *Group {
	return &Group{prefix: prefix}
}

func (g *Group) Handle(method string, path string, handler http.HandlerFunc) {
	g.routes = append(g.routes, route{
		Method:  method,
		Path:    g.prefix + path,
		Handler: handler,
	})
}

type Server struct {
	host   string
	port   string
	logger *slog.Logger
	mux    *http.ServeMux
}

func New(host string, port string, logger *slog.Logger) *Server {
	return &Server{
		host:   host,
		port:   port,
		logger: logger,
		mux:    http.NewServeMux(),
	}
}

func (s *Server) Register(method string, path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, withMethod(method, s.withMiddleware(handler)))
}

func (s *Server) RegisterGroup(group *Group) {
	for _, route := range group.routes {
		s.Register(route.Method, route.Path, route.Handler)
	}
}

func (s *Server) Address() string {
	return fmt.Sprintf("%s:%s", s.host, s.port)
}

func (s *Server) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    s.Address(),
		Handler: s.mux,
	}

	errCh := make(chan error, 1)

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		errCh <- server.Shutdown(shutdownCtx)
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return <-errCh
}

func (s *Server) withMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startedAt := time.Now()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)

		s.logger.Debug("http request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(startedAt).String(),
		)
	}
}

func withMethod(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{
				"error": fmt.Sprintf("method %s is not allowed for %s", r.Method, strings.TrimSpace(r.URL.Path)),
			})
			return
		}

		next(w, r)
	}
}

func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
