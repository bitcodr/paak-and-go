package trip

import (
	"context"
	"errors"
	"net/http"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/repository/impl"
)

type ITrip interface {
	List(ctx context.Context) ([]*model.Trip, error)
	Show(ctx context.Context, id int32) (*model.Trip, error)
	Store(ctx context.Context, trip *model.Trip) (*model.Trip, error)
}

type trip struct {
	repo impl.Trip
}

func InitService(_ context.Context, repository impl.Trip) ITrip {
	return &trip{
		repo: repository,
	}
}

func (t *trip) List(ctx context.Context) (trips []*model.Trip, err error) {
	trips, err = t.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return trips, err
}

func (t *trip) Show(ctx context.Context, id int32) (trip *model.Trip, err error) {
	trip, err = t.repo.Show(ctx, id)
	if err != nil {
		return nil, err
	}

	return trip, err
}

func (t *trip) Store(ctx context.Context, request *model.Trip) (trip *model.Trip, err error) {
	if request == nil {
		return nil, errors.New(http.StatusText(http.StatusNoContent))
	}

	trip, err = t.repo.Store(ctx, request)
	if err != nil {
		return nil, err
	}

	return trip, err
}
