package impl

import (
	"context"
	"io"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
)

type Trip interface {
	io.Closer
	List(ctx context.Context) ([]*model.Trip, error)
	Show(ctx context.Context, id int32) (*model.Trip, error)
	Store(ctx context.Context, trip *model.Trip) (*model.Trip, error)
}

type City interface {
	io.Closer
	List(ctx context.Context) ([]*model.City, error)
	Show(ctx context.Context, id int32) (*model.City, error)
	Store(ctx context.Context, city *model.City) (*model.City, error)
}
