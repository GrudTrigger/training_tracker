package exercises_http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/george/training-tracker/backend/internal/core/domain"
	core_errors "github.com/george/training-tracker/backend/internal/core/errors"
	core_httpserver "github.com/george/training-tracker/backend/internal/core/httpserver"
)

type Service interface {
	Create(ctx context.Context, name string, muscleGroup string) (domain.Exercise, error)
	List(ctx context.Context) ([]domain.Exercise, error)
}

type Handler struct {
	service Service
}

type createExerciseRequest struct {
	Name        string `json:"name"`
	MuscleGroup string `json:"muscleGroup"`
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(group *core_httpserver.Group) {
	group.Handle(http.MethodGet, "/exercises", h.List)
	group.Handle(http.MethodPost, "/exercises", h.Create)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	exercises, err := h.service.List(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}

	core_httpserver.WriteJSON(w, http.StatusOK, exercises)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var request createExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		core_httpserver.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	exercise, err := h.service.Create(r.Context(), request.Name, request.MuscleGroup)
	if err != nil {
		writeError(w, err)
		return
	}

	core_httpserver.WriteJSON(w, http.StatusCreated, exercise)
}

func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, core_errors.ErrValidation):
		core_httpserver.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	case errors.Is(err, core_errors.ErrConflict):
		core_httpserver.WriteJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
	default:
		core_httpserver.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
}
