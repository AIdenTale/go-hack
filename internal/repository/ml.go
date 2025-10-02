package repository

import (
	"context"
	"github.com/AIdenTale/go-hack.git/internal/model"
)

type MLRepository interface {
	// SavePrediction сохраняет результаты предсказания в БД
	SavePrediction(ctx context.Context, pred *model.MLPrediction) error
	
	// GetPredictions возвращает последние предсказания
	GetPredictions(ctx context.Context, limit int) ([]*model.MLPrediction, error)

	// GetLatestPrediction возвращает самую последнюю запись
	GetLatestPrediction(ctx context.Context) (*model.MLPrediction, error)
}