package main

import (
	"context"
	"log"

	"github.com/bitcodr/paak-and-go/internal/domain/service/trip"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/repository"
	"github.com/bitcodr/paak-and-go/internal/interfaces/transport"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	repo, err := repository.New(ctx, config.POSTGRES, &cfg.DB)
	if err != nil {
		log.Fatalln(err.Error())
	}

	tripService := trip.New(ctx, repo)

	transport.NewRest(ctx, &transport.Service{
		TripService: tripService,
	}, &cfg.Service)
}
