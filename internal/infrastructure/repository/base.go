package repository

import (
	"context"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/repository/postgres"
)

type Repository interface {
	List(ctx context.Context) ([]*model.Trip, error)
	Show(ctx context.Context, id int32) (*model.Trip, error)
	Store(ctx context.Context, trip *model.Trip) (*model.Trip, error)
}

func New(_ context.Context, repo Repository, config *config.DB) Repository {
	switch repo.(type) {
	case *postgres.Postgres:
		return &postgres.Postgres{Cfg: config}
	default:
		return nil
	}
}
