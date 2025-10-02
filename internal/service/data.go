package service

import (
	"context"

	"github.com/AIdenTale/go-hack.git/internal/model"
	"github.com/AIdenTale/go-hack.git/internal/repository"
)

type DataService struct {
	repo repository.DataRepository
}

func NewDataService(repo repository.DataRepository) *DataService {
	return &DataService{repo: repo}
}

func (s *DataService) GetAllData(ctx context.Context, seconds int64) (*model.DataResponse, error) {
	return s.repo.GetAllData(ctx, seconds)
}

func (s *DataService) GetFHRUpdates(ctx context.Context) (*model.DataResponse, error) {
	return s.repo.GetFHRUpdates(ctx)
}

func (s *DataService) GetUCUpdates(ctx context.Context) (*model.DataResponse, error) {
	return s.repo.GetUCUpdates(ctx)
}