package service

import (
	"context"
	"fmt"
	"time"

	"github.com/AIdenTale/go-hack.git/internal/model"
	"github.com/AIdenTale/go-hack.git/internal/repository"
)

type MLService struct {
	dataService *DataService
	mlClient    *MLClient
	mlRepo      repository.MLRepository
}

func NewMLService(
	dataService *DataService,
	mlClient *MLClient,
	mlRepo repository.MLRepository,
) *MLService {
	return &MLService{
		dataService: dataService,
		mlClient:    mlClient,
		mlRepo:      mlRepo,
	}
}

// ProcessLatestData получает последние данные и отправляет их в ML сервис
func (s *MLService) ProcessLatestData(ctx context.Context) error {
	// Получаем данные за последние 10 минут
	data, err := s.dataService.GetAllData(ctx, 600) // 600 секунд = 10 минут
	if err != nil {
		return fmt.Errorf("get data: %w", err)
	}

	if len(data.FHR) == 0 || len(data.UC) == 0 {
		return fmt.Errorf("no data available")
	}

	// Подготавливаем данные для ML сервиса
	mlData := &mlRequestData{
		FHRTime:   make([]float64, len(data.FHR)),
		FHRValues: make([]float64, len(data.FHR)),
		UCTime:    make([]float64, len(data.UC)),
		UCValues:  make([]float64, len(data.UC)),
	}

	// Преобразуем время в float64 (секунды с начала данных)
	startTime := data.FHR[0].Time
	for i, p := range data.FHR {
		mlData.FHRTime[i] = p.Time.Sub(startTime).Seconds()
		mlData.FHRValues[i] = p.D
	}
	for i, p := range data.UC {
		mlData.UCTime[i] = p.Time.Sub(startTime).Seconds()
		mlData.UCValues[i] = p.D
	}

	// Получаем предсказание
	predResp, err := s.mlClient.Predict(mlData)
	if err != nil {
		return fmt.Errorf("predict: %w", err)
	}

	// Получаем базовые метрики
	featResp, err := s.mlClient.GetFeatures(mlData)
	if err != nil {
		return fmt.Errorf("get features: %w", err)
	}

	// Создаем объект для сохранения
	prediction := &model.MLPrediction{
		TimeRecorded: time.Now().UTC(),
		Prediction:   predResp.Prediction,
		Probability:  predResp.Probability,
	}

	// Добавляем значения из top_features
	for _, f := range predResp.TopFeatures {
		switch f.Name {
		case "median_fhr":
			prediction.MedianFHRValue = f.Value
			prediction.MedianFHRImpact = f.Impact
		case "mean_fhr":
			prediction.MeanFHRValue = f.Value
			prediction.MeanFHRImpact = f.Impact
		case "cross_corr_fhr_uc":
			prediction.CrossCorrValue = f.Value
			prediction.CrossCorrImpact = f.Impact
		}
	}

	// Добавляем базовые метрики
	if len(featResp.Features) >= 8 {
		prediction.MeanFHR = featResp.Features[0]
		prediction.BaselineFHR = featResp.Features[1]
		prediction.MinFHR = featResp.Features[2]
		prediction.MaxFHR = featResp.Features[3]
		prediction.Accelerations = int(featResp.Features[4])
		prediction.Decelerations = int(featResp.Features[5])
		prediction.MeanUC = featResp.Features[6]
		prediction.MaxUC = featResp.Features[7]
	}

	// Сохраняем в БД
	if err := s.mlRepo.SavePrediction(ctx, prediction); err != nil {
		return fmt.Errorf("save prediction: %w", err)
	}

	return nil
}

// GetLatestPredictions возвращает последние предсказания
func (s *MLService) GetLatestPredictions(ctx context.Context, limit int) ([]*model.MLPrediction, error) {
	return s.mlRepo.GetPredictions(ctx, limit)
}

// GetLatestPrediction returns only the most recent prediction
func (s *MLService) GetLatestPrediction(ctx context.Context) (*model.MLPrediction, error) {
    return s.mlRepo.GetLatestPrediction(ctx)
}