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
	dataService *service.DataService,
	mlService *service.MLService) *App {
	return &App{
		Config:             cfg,
		Logger:             logger,
		Postgres:           pg,
		PregnantDatService: pregnantService,
		DataService:        dataService,
		MLService:          mlService,
	}
}

// AppDeps агрегирует зависимости приложения для внедрения через wire.
type App struct {
	Config *config.Config
	Logger *zap.Logger

	Postgres           *db.Postgres
	PregnantDatService *service.PregnantDatService
	DataService        *service.DataService
	MLService          *service.MLService
}

// InitializeApp инициализирует зависимости приложения через wire.
// provideMLClient создает клиент для ML сервиса
func provideMLClient(cfg *config.Config) *service.MLClient {
	return service.NewMLClient(cfg.ML.BaseURL)
}

// provideMLRepository создает ML репозиторий
func provideMLRepository(pg *db.Postgres) repository.MLRepository {
	return db.NewMLPostgresRepository(pg)
}

// Provide MLService
func provideMLService(
	dataService *service.DataService,
	mlClient *service.MLClient,
	mlRepo repository.MLRepository,
) *service.MLService {
	return service.NewMLService(dataService, mlClient, mlRepo)
}

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
		provideMLClient,
		provideMLRepository,
		provideMLService,
	)
	return nil, nil
}
