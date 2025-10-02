package main

import (
	"log"

	"github.com/AIdenTale/go-hack.git/internal/app"
	"github.com/AIdenTale/go-hack.git/internal/handlers"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	deps, err := app.InitializeApp("config/receiver.yaml")
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	e := echo.New()
	handlers.Register(deps, e)

	deps.Logger.Info("Starting server", zap.String("address", deps.Config.Echo.Address))
	if err := e.Start(deps.Config.Echo.Address); err != nil {
		deps.Logger.Fatal("server error", zap.Error(err))
	}
}
