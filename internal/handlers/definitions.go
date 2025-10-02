package handlers

import (
	"github.com/AIdenTale/go-hack.git/internal/app"
	"github.com/AIdenTale/go-hack.git/internal/handlers/views"
	"github.com/AIdenTale/go-hack.git/internal/handlers/views/bpm"
	"github.com/AIdenTale/go-hack.git/internal/handlers/views/trac"
	"github.com/AIdenTale/go-hack.git/internal/service"
	"github.com/labstack/echo/v4"
)

// Register регистрирует HTTP endpoint'ы для работы с bpm и trac.
// Принимает все зависимости приложения через *app.App.
func Register(app *app.App, mux *echo.Echo) {
	v1 := mux.Group("/api/v1")
	registerBPM(v1, app.PregnantDatService)
	registerTrac(v1, app.PregnantDatService)

	dataGroup := mux.Group("/data")
	registerData(dataGroup, app.DataService)
}

// registerBPM регистрирует endpoint'ы внутри группы /api/v1.
func registerBPM(mux *echo.Group, service *service.PregnantDatService) {
	handler := bpm.NewHandler(service)
	mux.POST("/bpm", handler.Bpm)
}

func registerTrac(mux *echo.Group, service *service.PregnantDatService) {
	handler := trac.NewHandler(service)
	mux.POST("/trac", handler.Trac)
}

func registerData(mux *echo.Group, service *service.DataService) {
	handler := views.NewDataHandler(service)
	mux.GET("/get_all", handler.GetAllData)
	mux.GET("/fhr/updates", handler.GetFHRUpdates)
	mux.GET("/uc/updates", handler.GetUCUpdates)
}
