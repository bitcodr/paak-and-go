package main

import (
	"context"
	"log"

	"github.com/bitcodr/paak-and-go/internal/domain/service/trip"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
	triprepo "github.com/bitcodr/paak-and-go/internal/infrastructure/repository/postgres/trip"
	"github.com/bitcodr/paak-and-go/internal/interfaces/transport"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	repo, err := triprepo.InitRepo(ctx, cfg.Connections[config.POSTGRES])
	if err != nil {
		log.Fatalln(err.Error())
	}

	tripService := trip.InitService(ctx, repo)

	transport.InitRest(ctx, &transport.Service{
		TripService: tripService,
	}, &cfg.Service)
}
