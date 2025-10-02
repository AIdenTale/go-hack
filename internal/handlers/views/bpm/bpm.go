package bpm

import (
	"github.com/AIdenTale/go-hack.git/internal/model"
	"github.com/AIdenTale/go-hack.git/internal/service"
	"github.com/labstack/echo/v4"
)

// Handler реализует HTTP-обработчик для bpm.
type Handler struct {
	service *service.PregnantDatService
}

// NewHandler создает новый Handler для bpm.
func NewHandler(service *service.PregnantDatService) *Handler {
	return &Handler{
		 service: service,
	}
}

// Bpm обрабатывает POST-запросы на /bpm.
func (h *Handler) Bpm(ctx echo.Context) error {
	var data []model.BPM
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(400, map[string]string{"error": err.Error()})
	}
	for _, d := range data {
		if err := h.service.InsertBPM(ctx.Request().Context(), d); err != nil {
			return ctx.JSON(500, map[string]string{"error": err.Error()})
		}
	}
	return ctx.JSON(200, map[string]string{"status": "ok"})
}
