package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
)

type TripRepo struct {
	mock.Mock
}

func (m *TripRepo) List(ctx context.Context) ([]*model.Trip, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.Trip), args.Error(1)
}

func (m *TripRepo) Show(ctx context.Context, id int32) (*model.Trip, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Trip), args.Error(1)
}

func (m *TripRepo) Store(ctx context.Context, trip *model.Trip) (*model.Trip, error) {
	args := m.Called(ctx, trip)
	return args.Get(0).(*model.Trip), args.Error(1)
}

func (*TripRepo) Close() error {
	//todo implement
	return nil
}
