package views

import (
	"net/http"
	"strconv"

	"github.com/AIdenTale/go-hack.git/internal/service"
	"github.com/labstack/echo/v4"
)

type DataHandler struct {
	service *service.DataService
}

func NewDataHandler(service *service.DataService) *DataHandler {
	return &DataHandler{service: service}
}

// GetAllData godoc
// @Summary Get all FHR and UC data for a time range
// @Description Get all FHR and UC data for the specified number of seconds back from now
// @Tags data
// @Accept json
// @Produce json
// @Param seconds query int true "Number of seconds to look back"
// @Success 200 {object} model.DataResponse
// @Router /data/get_all [get]
func (h *DataHandler) GetAllData(c echo.Context) error {
	seconds, err := strconv.ParseInt(c.QueryParam("seconds"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid seconds parameter")
	}

	data, err := h.service.GetAllData(c.Request().Context(), seconds)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

// GetFHRUpdates godoc
// @Summary Get new FHR data updates
// @Description Get up to 10000 new FHR records that haven't been retrieved yet
// @Tags data
// @Produce json
// @Success 200 {object} model.DataResponse
// @Router /data/fhr/updates [get]
func (h *DataHandler) GetFHRUpdates(c echo.Context) error {
	data, err := h.service.GetFHRUpdates(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

// GetUCUpdates godoc
// @Summary Get new UC data updates
// @Description Get up to 10000 new UC records that haven't been retrieved yet
// @Tags data
// @Produce json
// @Success 200 {object} model.DataResponse
// @Router /data/uc/updates [get]
func (h *DataHandler) GetUCUpdates(c echo.Context) error {
	data, err := h.service.GetUCUpdates(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}