package repository

import (
	"context"
	"errors"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/repository/postgres/trip"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/helper"
)

type Repository interface {
	List(ctx context.Context) ([]*model.Trip, error)
	Show(ctx context.Context, id int32) (*model.Trip, error)
	Store(ctx context.Context, trip *model.Trip) (*model.Trip, error)
}

func New(ctx context.Context, connectionName string, cfg *config.DB) (Repository, error) {
	if cfg == nil {
		return nil, errors.New("config is empty")
	}

	switch connectionName {
	case config.POSTGRES:

		postgresConfig := cfg.Connections[config.POSTGRES]

		connection, err := helper.NewPostgres(ctx, postgresConfig)
		if err != nil {
			return nil, err
		}

		return &trip.Trip{
			Cfg:  postgresConfig,
			Conn: connection,
		}, nil

	default:
		return nil, errors.New("connection does not exist")
	}
}
