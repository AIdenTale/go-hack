package service

import (
	"context"
	"github.com/AIdenTale/go-hack.git/internal/model"
	"github.com/AIdenTale/go-hack.git/internal/repository"
)

type Front struct {
	repo repository.PregnantDatRepository
}

func NewFront(repo repository.PregnantDatRepository) *Front {
	return &Front{repo: repo}
}

func (s *Front) GetAll(ctx context.Context, bpm model.BPM) error {
	return s.repo.InsertBPM(ctx, bpm)
}

