package trip

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
)

//MockRepo will satisfy the TripRepo interface for testing purpose
//we don't wanna have an actual insert in db
//to use test cases in CI it is best to use mocks
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) List(ctx context.Context) ([]*model.Trip, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.Trip), args.Error(1)
}

func (m *MockRepo) Show(ctx context.Context, id int32) (*model.Trip, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Trip), args.Error(1)
}

func (m *MockRepo) Store(ctx context.Context, trip *model.Trip) (*model.Trip, error) {
	args := m.Called(ctx, trip)
	return args.Get(0).(*model.Trip), args.Error(1)
}

func (*MockRepo) Close() error {
	//todo implement
	return nil
}
