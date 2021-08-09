package trip

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/helper"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/repository/impl"
)

type trip struct {
	cfg  *config.Connection
	conn *pgxpool.Pool
}

func InitRepo(ctx context.Context, cfg *config.Connection) (impl.Trip, error) {
	if cfg == nil {
		return nil, errors.New("config is empty")
	}

	connection, err := helper.NewPostgres(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &trip{
		cfg:  cfg,
		conn: connection,
	}, nil
}

func (p *trip) List(ctx context.Context) ([]*model.Trip, error) {
	defer p.conn.Close()

	return nil, nil
}

func (p *trip) Show(ctx context.Context, id int32) (*model.Trip, error) {
	defer p.conn.Close()

	return nil, nil
}

func (p *trip) Store(ctx context.Context, trip *model.Trip) (*model.Trip, error) {
	defer p.conn.Close()

	return nil, nil
}
