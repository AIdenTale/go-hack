package db

import (
	"context"
	"github.com/AIdenTale/go-hack.git/internal/model"
)

type MLPostgresRepository struct {
	pg *Postgres
}

func NewMLPostgresRepository(pool *Postgres) *MLPostgresRepository {
	return &MLPostgresRepository{pg: pool}
}

func (r *MLPostgresRepository) SavePrediction(ctx context.Context, pred *model.MLPrediction) error {
	_, err := r.pg.Pool.Exec(ctx, `
		INSERT INTO ml_predictions (
			time_recorded,
			prediction,
			probability,
			median_fhr_value,
			median_fhr_impact,
			mean_fhr_value,
			mean_fhr_impact,
			cross_corr_fhr_uc_value,
			cross_corr_fhr_uc_impact,
			mean_fhr,
			baseline_fhr,
			min_fhr,
			max_fhr,
			accelerations,
			decelerations,
			mean_uc,
			max_uc
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
		)`,
		pred.TimeRecorded,
		pred.Prediction,
		pred.Probability,
		pred.MedianFHRValue,
		pred.MedianFHRImpact,
		pred.MeanFHRValue,
		pred.MeanFHRImpact,
		pred.CrossCorrValue,
		pred.CrossCorrImpact,
		pred.MeanFHR,
		pred.BaselineFHR,
		pred.MinFHR,
		pred.MaxFHR,
		pred.Accelerations,
		pred.Decelerations,
		pred.MeanUC,
		pred.MaxUC,
	)
	return err
}

func (r *MLPostgresRepository) GetPredictions(ctx context.Context, limit int) ([]*model.MLPrediction, error) {
	rows, err := r.pg.Pool.Query(ctx, `
		SELECT 
			time_recorded,
			prediction,
			probability,
			median_fhr_value,
			median_fhr_impact,
			mean_fhr_value,
			mean_fhr_impact,
			cross_corr_fhr_uc_value,
			cross_corr_fhr_uc_impact,
			mean_fhr,
			baseline_fhr,
			min_fhr,
			max_fhr,
			accelerations,
			decelerations,
			mean_uc,
			max_uc
		FROM ml_predictions
		ORDER BY time_recorded DESC
		LIMIT $1`,
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var predictions []*model.MLPrediction
	for rows.Next() {
		pred := &model.MLPrediction{}
		err := rows.Scan(
			&pred.TimeRecorded,
			&pred.Prediction,
			&pred.Probability,
			&pred.MedianFHRValue,
			&pred.MedianFHRImpact,
			&pred.MeanFHRValue,
			&pred.MeanFHRImpact,
			&pred.CrossCorrValue,
			&pred.CrossCorrImpact,
			&pred.MeanFHR,
			&pred.BaselineFHR,
			&pred.MinFHR,
			&pred.MaxFHR,
			&pred.Accelerations,
			&pred.Decelerations,
			&pred.MeanUC,
			&pred.MaxUC,
		)
		if err != nil {
			return nil, err
		}
		predictions = append(predictions, pred)
	}
	return predictions, nil
}

func (r *MLPostgresRepository) GetLatestPrediction(ctx context.Context) (*model.MLPrediction, error) {
    row := r.pg.Pool.QueryRow(ctx, `
        SELECT 
            time_recorded,
            prediction,
            probability,
            median_fhr_value,
            median_fhr_impact,
            mean_fhr_value,
            mean_fhr_impact,
            cross_corr_fhr_uc_value,
            cross_corr_fhr_uc_impact,
            mean_fhr,
            baseline_fhr,
            min_fhr,
            max_fhr,
            accelerations,
            decelerations,
            mean_uc,
            max_uc
        FROM ml_predictions
        ORDER BY time_recorded DESC
        LIMIT 1`)

    pred := &model.MLPrediction{}
    if err := row.Scan(
        &pred.TimeRecorded,
        &pred.Prediction,
        &pred.Probability,
        &pred.MedianFHRValue,
        &pred.MedianFHRImpact,
        &pred.MeanFHRValue,
        &pred.MeanFHRImpact,
        &pred.CrossCorrValue,
        &pred.CrossCorrImpact,
        &pred.MeanFHR,
        &pred.BaselineFHR,
        &pred.MinFHR,
        &pred.MaxFHR,
        &pred.Accelerations,
        &pred.Decelerations,
        &pred.MeanUC,
        &pred.MaxUC,
    ); err != nil {
        return nil, err
    }
    return pred, nil
}