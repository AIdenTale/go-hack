package trac

import (
	"log"

	"github.com/AIdenTale/go-hack.git/internal/model"
	"github.com/AIdenTale/go-hack.git/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.PregnantDatService
}

func NewHandler(service *service.PregnantDatService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Trac(ctx echo.Context) error {
	var data []model.Trac
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(400, map[string]string{"error": err.Error()})
	}
	for _, d := range data {
		if err := h.service.InsertTrac(ctx.Request().Context(), d); err != nil {
			log.Panicln(err)
			return ctx.JSON(500, map[string]string{"error": err.Error()})
		}
	}
	return ctx.JSON(200, map[string]string{"status": "ok"})
}
