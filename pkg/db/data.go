package db

import (
	"context"
	"time"

	"github.com/AIdenTale/go-hack.git/internal/model"
)

type DataPostgresRepository struct {
	pg *Postgres
}

func NewDataPostgresRepository(pool *Postgres) *DataPostgresRepository {
	return &DataPostgresRepository{pg: pool}
}

func (r *DataPostgresRepository) GetAllData(ctx context.Context, seconds int64) (*model.DataResponse, error) {
	tx, err := r.pg.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	timeFrom := time.Now().Add(-time.Duration(seconds) * time.Second)
	
	// Get and mark FHR data
	fhrRows, err := tx.Query(ctx, `
		WITH records AS (
			SELECT time_recorded, value 
			FROM fhr_records 
			WHERE time_recorded >= $1
			FOR UPDATE
		)
		UPDATE fhr_records fr
		SET is_retrieved = true
		FROM records r
		WHERE fr.time_recorded = r.time_recorded
		RETURNING r.time_recorded, r.value`, timeFrom)
	if err != nil {
		return nil, err
	}
	defer fhrRows.Close()

	var fhrData []model.DataPoint
	for fhrRows.Next() {
		var dp model.DataPoint
		if err := fhrRows.Scan(&dp.Time, &dp.D); err != nil {
			return nil, err
		}
		fhrData = append(fhrData, dp)
	}

	// Get and mark UC data
	ucRows, err := tx.Query(ctx, `
		WITH records AS (
			SELECT time_recorded, value 
			FROM uc_records 
			WHERE time_recorded >= $1
			FOR UPDATE
		)
		UPDATE uc_records ur
		SET is_retrieved = true
		FROM records r
		WHERE ur.time_recorded = r.time_recorded
		RETURNING r.time_recorded, r.value`, timeFrom)
	if err != nil {
		return nil, err
	}
	defer ucRows.Close()

	var ucData []model.DataPoint
	for ucRows.Next() {
		var dp model.DataPoint
		if err := ucRows.Scan(&dp.Time, &dp.D); err != nil {
			return nil, err
		}
		ucData = append(ucData, dp)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &model.DataResponse{
		FHR: fhrData,
		UC:  ucData,
		LastChunk: true, // всегда true для GetAllData
	}, nil
}

func (r *DataPostgresRepository) GetFHRUpdates(ctx context.Context) (*model.DataResponse, error) {
	tx, err := r.pg.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Get new FHR records and mark them as retrieved
	rows, err := tx.Query(ctx, `
		WITH new_records AS (
			SELECT id, time_recorded, value 
			FROM fhr_records 
			WHERE NOT is_retrieved
			ORDER BY time_recorded
			LIMIT 10000
			FOR UPDATE
		)
		UPDATE fhr_records fr
		SET is_retrieved = true
		FROM new_records nr
		WHERE fr.id = nr.id
		RETURNING nr.time_recorded, nr.value`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []model.DataPoint
	for rows.Next() {
		var dp model.DataPoint
		if err := rows.Scan(&dp.Time, &dp.D); err != nil {
			return nil, err
		}
		data = append(data, dp)
	}

	// Check if there are more unretrieved records
	var remainingCount int
	err = tx.QueryRow(ctx, `
		SELECT COUNT(*) 
		FROM fhr_records 
		WHERE NOT is_retrieved`).Scan(&remainingCount)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &model.DataResponse{
		FHR:       data,
		LastChunk: remainingCount == 0,
	}, nil
}

func (r *DataPostgresRepository) GetUCUpdates(ctx context.Context) (*model.DataResponse, error) {
	tx, err := r.pg.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Get new UC records and mark them as retrieved
	rows, err := tx.Query(ctx, `
		WITH new_records AS (
			SELECT id, time_recorded, value 
			FROM uc_records 
			WHERE NOT is_retrieved
			ORDER BY time_recorded
			LIMIT 10000
			FOR UPDATE
		)
		UPDATE uc_records ur
		SET is_retrieved = true
		FROM new_records nr
		WHERE ur.id = nr.id
		RETURNING nr.time_recorded, nr.value`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []model.DataPoint
	for rows.Next() {
		var dp model.DataPoint
		if err := rows.Scan(&dp.Time, &dp.D); err != nil {
			return nil, err
		}
		data = append(data, dp)
	}

	// Check if there are more unretrieved records
	var remainingCount int
	err = tx.QueryRow(ctx, `
		SELECT COUNT(*) 
		FROM uc_records 
		WHERE NOT is_retrieved`).Scan(&remainingCount)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &model.DataResponse{
		UC:        data,
		LastChunk: remainingCount == 0,
	}, nil
}