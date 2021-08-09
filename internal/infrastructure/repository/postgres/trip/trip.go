package trip

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
)

type Trip struct {
	Cfg  *config.Connection
	Conn *pgxpool.Pool
}

func (p *Trip) List(ctx context.Context) ([]*model.Trip, error) {
	defer p.Conn.Close()

	p.Conn.Query(ctx, "select `id`, ``")
}

func (p *Trip) Show(ctx context.Context, id int32) (*model.Trip, error) {
	defer p.Conn.Close()

}

func (p *Trip) Store(ctx context.Context, trip *model.Trip) (*model.Trip, error) {
	defer p.Conn.Close()
}
