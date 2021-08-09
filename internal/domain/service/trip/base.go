package trip

import (
	"context"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/repository"
)

type ITrip interface {
	List(ctx context.Context) ([]*model.Trip, error)
	Show(ctx context.Context, id int32) (*model.Trip, error)
	Store(ctx context.Context, trip *model.Trip) (*model.Trip, error)
}

func New(_ context.Context, repository repository.Repository) ITrip {
	return &trip{
		repo: repository,
	}
}
