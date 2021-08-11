package impl

import (
	"context"
	"io"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
)

type TripRepo interface {
	io.Closer //to close db connection in the service termination
	List(ctx context.Context) ([]*model.Trip, error)
	Show(ctx context.Context, id int32) (*model.Trip, error)
	Store(ctx context.Context, trip *model.Trip) (*model.Trip, error)
}


//add other entity interfaces in here