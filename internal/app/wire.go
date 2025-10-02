//go:build wireinject
// +build wireinject

package app

import (
	"github.com/AIdenTale/go-hack.git/internal/app/config"
	"github.com/AIdenTale/go-hack.git/internal/repository"
	"github.com/AIdenTale/go-hack.git/internal/service"
	"github.com/AIdenTale/go-hack.git/pkg/db"
	"github.com/AIdenTale/go-hack.git/pkg/logger"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func newApp(cfg *config.Config, logger *zap.Logger, pg *db.Postgres, 
	pregnantService *service.PregnantDatService,
	dataService *service.DataService) *App {
	return &App{
		Config:             cfg,
		Logger:             logger,
		Postgres:           pg,
		PregnantDatService: pregnantService,
		DataService:        dataService,
	}
}

// AppDeps агрегирует зависимости приложения для внедрения через wire.
type App struct {
	Config *config.Config
	Logger *zap.Logger

	Postgres           *db.Postgres
	PregnantDatService *service.PregnantDatService
	DataService        *service.DataService
}

// InitializeApp инициализирует зависимости приложения через wire.
func InitializeApp(configPath string) (*App, error) {
	wire.Build(
		config.LoadConfig,
		newApp,
		logger.New,
		db.New,
		db.NewPregnantDatPostgresRepository,
		db.NewDataPostgresRepository,
		wire.Bind(new(repository.PregnantDatRepository), new(*db.PregnantDatPostgresRepository)),
		wire.Bind(new(repository.DataRepository), new(*db.DataPostgresRepository)),
		service.NewPregnantDatService,
		service.NewDataService,
	)
	return nil, nil
}
