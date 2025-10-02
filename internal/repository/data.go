package repository

import (
	"context"

	"github.com/AIdenTale/go-hack.git/internal/model"
)

type DataRepository interface {
	GetAllData(ctx context.Context, seconds int64) (*model.DataResponse, error)
	GetFHRUpdates(ctx context.Context) (*model.DataResponse, error)
	GetUCUpdates(ctx context.Context) (*model.DataResponse, error)
}