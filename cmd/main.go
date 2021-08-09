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

	repo := repository.New(ctx, nil, &cfg.DB)
	if repo == nil {
		log.Fatalln("repo is not exist")
	}

	tripService := trip.New(ctx, repo)

	transport.NewRest(ctx, &transport.Service{
		TripService: tripService,
	}, &cfg.Service)
}
