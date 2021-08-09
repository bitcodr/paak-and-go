package city

import (
	"context"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type City struct {
	Cfg  *config.Connection
	Conn *pgxpool.Pool
}

func (c *City) List(ctx context.Context) ([]*model.Trip, error) {
	//todo implement

	return nil, nil
}

func (c *City) Show(ctx context.Context, id int32) (*model.Trip, error) {
	//todo implement

	return nil, nil
}

func (c *City) Store(ctx context.Context, city *model.Trip) (*model.Trip, error) {
	//todo implement

	return nil, nil
}
