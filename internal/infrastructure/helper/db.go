package helper

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
)

//NewPostgres start connection for postgres
func NewPostgres(ctx context.Context, cfg *config.Connection) (*pgxpool.Pool, error) {
	if cfg == nil {
		return nil, errors.New("config is empty")
	}

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host,
		strconv.Itoa(cfg.Port), cfg.Name, cfg.Ssl)

	conn, err := pgxpool.Connect(ctx, connection)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
