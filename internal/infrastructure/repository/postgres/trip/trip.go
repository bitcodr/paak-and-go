package trip

import (
	"context"
	"errors"
	"net/http"

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

func InitRepo(ctx context.Context, cfg *config.Connection) (impl.TripRepo, error) {
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

func (t *trip) Close() error {
	if t.conn != nil {
		t.conn.Close()
	}

	return nil
}

func (t *trip) List(ctx context.Context) (trips []*model.Trip, err error) {
	//todo add OFFSET, LIMIT for pagination

	query, err := t.conn.Query(ctx, `SELECT 
				trips.id, trips.dates, trips.price, origin.name, destination.name
				FROM trips 
				INNER JOIN cities AS origin ON trips.origin_id = origin.id
				INNER JOIN cities AS destination ON trips.destination_id = destination.id
				ORDER BY trips.created_at DESC`)

	if err != nil {
		return nil, err
	}

	defer query.Close()

	if !query.Next() {
		return nil, errors.New(http.StatusText(http.StatusNoContent))
	}

	for query.Next() {

		var trip model.Trip
		var origin, destination model.City

		err := query.Scan(&trip.ID, &trip.Dates, &trip.Price, &origin.Name, &destination.Name)
		if err != nil {
			return nil, err
		}

		trip.Origin = &origin
		trip.Destination = &destination

		trips = append(trips, &trip)
	}

	return trips, nil
}

func (t *trip) Show(ctx context.Context, id int32) (*model.Trip, error) {
	row := t.conn.QueryRow(ctx, `SELECT 
				trips.id, trips.dates, trips.price, origin.name, destination.name
				FROM trips 
				INNER JOIN cities AS origin ON trips.origin_id = origin.id
				INNER JOIN cities AS destination ON trips.destination_id = destination.id
				WHERE trips.id = $1`, id)

	var trip model.Trip

	var origin, destination model.City

	if err := row.Scan(&trip.ID, &trip.Dates, &trip.Price, &origin.Name, &destination.Name); err != nil {
		return nil, err
	}

	trip.Origin = &origin
	trip.Destination = &destination

	return &trip, nil
}

func (t *trip) Store(ctx context.Context, trip *model.Trip) (*model.Trip, error) {
	var tripId int32

	err := t.conn.QueryRow(ctx, `INSERT INTO trips 
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
