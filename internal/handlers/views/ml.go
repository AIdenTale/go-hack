package views

import (
	"net/http"

	"github.com/AIdenTale/go-hack.git/internal/service"
	"github.com/labstack/echo/v4"
)

type MLHandler struct {
	service *service.MLService
}

func NewMLHandler(service *service.MLService) *MLHandler {
	return &MLHandler{service: service}
}

// GetPredicts godoc
// @Summary Get most recent ML prediction
// @Description Returns the latest prediction from ML service stored in DB
// @Tags ml
// @Produce json
// @Success 200 {object} model.MLPrediction
// @Router /data/predicts [get]
func (h *MLHandler) GetPredicts(c echo.Context) error {
    prediction, err := h.service.GetLatestPrediction(c.Request().Context())
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, prediction)
}