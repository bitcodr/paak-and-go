package repository

import (
	"context"

	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
)

type Repository interface {
	List()
	Show()
	Store()
}


func New(_ context.Context, config *config.Config) Repository{

}
