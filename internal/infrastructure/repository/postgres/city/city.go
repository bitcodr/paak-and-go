package city

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/helper"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/repository/impl"
)

type city struct {
	cfg  *config.Connection
	conn *pgxpool.Pool
}

func InitRepo(ctx context.Context, cfg *config.Connection) (impl.CityRepo, error) {
	if cfg == nil {
		return nil, errors.New("config is empty")
	}

	connection, err := helper.NewPostgres(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &city{
		cfg:  cfg,
		conn: connection,
	}, nil
}

func (c *city) Close() error {
	if c.conn != nil {
		c.conn.Close()
	}

	return nil
}

func (c *city) List(ctx context.Context) ([]*model.City, error) {
	//todo impl

	return nil, nil
}

func (c *city) Show(ctx context.Context, id int32) (*model.City, error) {
	//todo impl

	return nil, nil
}

func (c *city) Store(ctx context.Context, city *model.City) (*model.City, error) {
	//todo impl

	return nil, nil
}
