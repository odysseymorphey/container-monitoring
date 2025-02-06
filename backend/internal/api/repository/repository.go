package repository

import (
	"container-monitoring/internal/api/models"
	"container-monitoring/internal/api/storage"
	"context"
)

type Repository interface {
	GetStatuses(ctx context.Context) ([]models.PingStatus, error)
	AddStatus(
		ctx context.Context,
		status *models.PingStatus,
	) error
	Close() error
}

type repository struct {
	storage *storage.PostgresDB
}

func NewRepository(storage *storage.PostgresDB) Repository {
	return &repository{storage: storage}
}

func (r *repository) GetStatuses(ctx context.Context) ([]models.PingStatus, error) {
	return r.storage.GetStatuses(ctx)
}

func (r *repository) AddStatus(ctx context.Context, status *models.PingStatus) error {
	return r.storage.AddStatus(ctx, status)
}

func (r *repository) Close() error {
	return r.storage.Close()
}
