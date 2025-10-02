// Реализация PregnantDatRepository для PostgreSQL.
package db

import (
	"context"

	"github.com/AIdenTale/go-hack.git/internal/model"
	"github.com/AIdenTale/go-hack.git/internal/repository"
)

type PregnantDatPostgresRepository struct {
	pg *Postgres
}

func NewPregnantDatPostgresRepository(pool *Postgres) *PregnantDatPostgresRepository {
	return &PregnantDatPostgresRepository{pg: pool}
}

func (r *PregnantDatPostgresRepository) InsertBPM(ctx context.Context, bpm model.BPM) error {
	_, err := r.pg.Pool.Exec(ctx, `
		INSERT INTO fhr_records (time_recorded, value) 
		VALUES (NOW(), $1)`, 
		bpm.BPM)
	return err
}

func (r *PregnantDatPostgresRepository) InsertTrac(ctx context.Context, trac model.Trac) error {
	_, err := r.pg.Pool.Exec(ctx, `
		INSERT INTO uc_records (time_recorded, value) 
		VALUES (NOW(), $1)`, 
		trac.Trac)
	return err
}

var _ repository.PregnantDatRepository = (*PregnantDatPostgresRepository)(nil)
