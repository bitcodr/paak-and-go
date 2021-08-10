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

func (p *trip) List(ctx context.Context, offset, limit int) (trips []*model.Trip, err error) {
	defer p.conn.Close()

	query, err := p.conn.Query(ctx, `SELECT 
				id, origin_id, destination_id, dates, price 
				FROM trips 
				INNER JOIN city AS origin ON trips.origin_id = origin.id
				INNER JOIN city AS destination ON trips.destination_id = destination.id
				ORDER BY trips.created_at DESC OFFSET $1 LIMIT $2`, offset, limit)

	if err != nil {
		return nil, err
	}

	defer query.Close()

	for query.Next() {
		var trip *model.Trip
		if err := query.Scan(&trip); err != nil {
			return nil, err
		}

		trips = append(trips, trip)
	}

	return trips, nil
}

func (p *trip) Show(ctx context.Context, id int32) (trip *model.Trip, err error) {
	defer p.conn.Close()

	row := p.conn.QueryRow(ctx, `SELECT 
				id, origin_id, destination_id, dates, price 
				FROM trips 
				INNER JOIN city AS origin ON trips.origin_id = origin.id
				INNER JOIN city AS destination ON trips.destination_id = destination.id
				WHERE trips.id = $1`, id)

	if err := row.Scan(&trip); err != nil {
		return nil, err
	}

	return trip, nil
}

func (p *trip) Store(ctx context.Context, trip *model.Trip) (*model.Trip, error) {
	defer p.conn.Close()

	var tripId int32

	err := p.conn.QueryRow(ctx, `INSERT INTO trips 
						(origin_id, destination_id, dates, price)
						VALUES ($1, $2, $3, $4)
						RETURNING id`,
		trip.Origin.ID, trip.Destination.ID, trip.Dates, trip.Price).Scan(&tripId)

	if err != nil {
		return nil, err
	}

	trip.ID = tripId

	return trip, nil
}
