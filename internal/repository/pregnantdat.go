// PregnantDatRepository определяет интерфейс для работы с bpm и trac.
package repository

import (
	"context"
	"github.com/AIdenTale/go-hack.git/internal/model"
)

type PregnantDatRepository interface {
	InsertBPM(ctx context.Context, bpm model.BPM) error
	InsertTrac(ctx context.Context, trac model.Trac) error
}
