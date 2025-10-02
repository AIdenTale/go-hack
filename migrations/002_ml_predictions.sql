CREATE TABLE IF NOT EXISTS ml_predictions (
    id SERIAL PRIMARY KEY,
    time_recorded TIMESTAMP WITH TIME ZONE NOT NULL,
    
    -- предсказание
    prediction INTEGER NOT NULL,           -- 0 = "норма", 1 = "гипоксия"
    probability DOUBLE PRECISION NOT NULL, -- вероятность гипоксии
    
    -- значения топ-признаков из /predict
    median_fhr_value DOUBLE PRECISION NOT NULL,
    median_fhr_impact DOUBLE PRECISION NOT NULL,
    
    mean_fhr_value DOUBLE PRECISION NOT NULL,
    mean_fhr_impact DOUBLE PRECISION NOT NULL,
    
    cross_corr_fhr_uc_value DOUBLE PRECISION NOT NULL,
    cross_corr_fhr_uc_impact DOUBLE PRECISION NOT NULL,
    
    -- значения базовых метрик из /features
    mean_fhr DOUBLE PRECISION NOT NULL,
    baseline_fhr DOUBLE PRECISION NOT NULL,
    min_fhr DOUBLE PRECISION NOT NULL,
    max_fhr DOUBLE PRECISION NOT NULL,
    accelerations INTEGER NOT NULL,
    decelerations INTEGER NOT NULL,
    mean_uc DOUBLE PRECISION NOT NULL,
    max_uc DOUBLE PRECISION NOT NULL,
    
    -- метаданные
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_ml_predictions_time ON ml_predictions(time_recorded);