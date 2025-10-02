// PregnantDatService агрегирует бизнес-логику для bpm и trac.
package service

import (
	"context"
	"github.com/AIdenTale/go-hack.git/internal/model"
	"github.com/AIdenTale/go-hack.git/internal/repository"
)

type PregnantDatService struct {
	repo repository.PregnantDatRepository
}

func NewPregnantDatService(repo repository.PregnantDatRepository) *PregnantDatService {
	return &PregnantDatService{repo: repo}
}

func (s *PregnantDatService) InsertBPM(ctx context.Context, bpm model.BPM) error {
	return s.repo.InsertBPM(ctx, bpm)
}

func (s *PregnantDatService) InsertTrac(ctx context.Context, trac model.Trac) error {
	return s.repo.InsertTrac(ctx, trac)
}
