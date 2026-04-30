package statistics_http

import (
	"context"
	"net/http"

	"github.com/george/training-tracker/backend/internal/core/domain"
	core_httpserver "github.com/george/training-tracker/backend/internal/core/httpserver"
)

type Service interface {
	GetOverview(ctx context.Context) (domain.StatisticsOverview, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(group *core_httpserver.Group) {
	group.Handle(http.MethodGet, "/statistics/overview", h.GetOverview)
}

func (h *Handler) GetOverview(w http.ResponseWriter, r *http.Request) {
	overview, err := h.service.GetOverview(r.Context())
	if err != nil {
		core_httpserver.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	core_httpserver.WriteJSON(w, http.StatusOK, overview)
}
