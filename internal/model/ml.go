package model

import "time"

// MLFeature представляет один важный признак из ML-модели
type MLFeature struct {
	Name   string  `json:"name"`
	Label  string  `json:"label"`
	Value  float64 `json:"value"`
	Impact float64 `json:"impact"`
}

// MLPredictResponse представляет ответ от /predict эндпоинта
type MLPredictResponse struct {
	Prediction   int         `json:"prediction"`
	Probability  float64     `json:"probability"`
	TopFeatures  []MLFeature `json:"top_features"`
}

// MLFeaturesResponse представляет ответ от /features эндпоинта
type MLFeaturesResponse struct {
	Features     []float64            `json:"features"`
	Descriptions map[string]string    `json:"descriptions"`
}

// MLPrediction объединяет все данные для сохранения в БД
type MLPrediction struct {
    TimeRecorded     time.Time  `json:"time_recorded"`
	Prediction       int     `json:"prediction"`
	Probability      float64 `json:"probability"`

	// Топ-признаки из /predict
	MedianFHRValue      float64 `json:"median_fhr_value"`
	MedianFHRImpact     float64 `json:"median_fhr_impact"`
	MeanFHRValue        float64 `json:"mean_fhr_value"`
	MeanFHRImpact       float64 `json:"mean_fhr_impact"`
	CrossCorrValue      float64 `json:"cross_corr_fhr_uc_value"`
	CrossCorrImpact     float64 `json:"cross_corr_fhr_uc_impact"`

	// Базовые метрики из /features
	MeanFHR       float64 `json:"mean_fhr"`
	BaselineFHR   float64 `json:"baseline_fhr"`
	MinFHR        float64 `json:"min_fhr"`
	MaxFHR        float64 `json:"max_fhr"`
	Accelerations int     `json:"accelerations"`
	Decelerations int     `json:"decelerations"`
	MeanUC        float64 `json:"mean_uc"`
	MaxUC         float64 `json:"max_uc"`
}