package workouts_http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/george/training-tracker/backend/internal/core/domain"
	core_errors "github.com/george/training-tracker/backend/internal/core/errors"
	core_httpserver "github.com/george/training-tracker/backend/internal/core/httpserver"
	workouts_service "github.com/george/training-tracker/backend/internal/features/workouts/service"
)

type Service interface {
	Create(ctx context.Context, performedAt time.Time, note string, sets []workouts_service.CreateSetInput) (domain.Workout, error)
	List(ctx context.Context) ([]domain.Workout, error)
}

type createWorkoutRequest struct {
	PerformedAt string `json:"performedAt"`
	Note        string `json:"note"`
	Sets        []struct {
		ExerciseID string  `json:"exerciseId"`
		Reps       int     `json:"reps"`
		WeightKg   float64 `json:"weightKg"`
		SetOrder   int     `json:"setOrder"`
	} `json:"sets"`
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(group *core_httpserver.Group) {
	group.Handle(http.MethodGet, "/workouts", h.List)
	group.Handle(http.MethodPost, "/workouts", h.Create)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	workouts, err := h.service.List(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}

	core_httpserver.WriteJSON(w, http.StatusOK, workouts)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var request createWorkoutRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		core_httpserver.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	performedAt, err := time.Parse(time.RFC3339, request.PerformedAt)
	if err != nil {
		core_httpserver.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "performedAt must be RFC3339"})
		return
	}

	sets := make([]workouts_service.CreateSetInput, 0, len(request.Sets))
	for _, set := range request.Sets {
		sets = append(sets, workouts_service.CreateSetInput{
			ExerciseID: set.ExerciseID,
			Reps:       set.Reps,
			WeightKg:   set.WeightKg,
			SetOrder:   set.SetOrder,
		})
	}

	workout, err := h.service.Create(r.Context(), performedAt, request.Note, sets)
	if err != nil {
		writeError(w, err)
		return
	}

	core_httpserver.WriteJSON(w, http.StatusCreated, workout)
}

func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, core_errors.ErrValidation):
		core_httpserver.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	default:
		core_httpserver.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
}
