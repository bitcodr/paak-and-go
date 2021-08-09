package postgres

import (
	"context"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
)

type Postgres struct {
	Cfg *config.DB
}

func (p *Postgres) List(ctx context.Context) ([]*model.Trip, error) {
	panic("implement me")
}

func (p *Postgres) Show(ctx context.Context, id int32) (*model.Trip, error) {
	panic("implement me")
}

func (p *Postgres) Store(ctx context.Context, trip *model.Trip) (*model.Trip, error) {
	panic("implement me")
}

