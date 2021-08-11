package main

import (
	"context"
	"log"

	"github.com/bitcodr/paak-and-go/internal/domain/service/trip"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
	triprepo "github.com/bitcodr/paak-and-go/internal/infrastructure/repository/postgres/trip"
	"github.com/bitcodr/paak-and-go/internal/interfaces/transport"
)

/*
main to know about how to run project you can check README.md
The structure of project is Hexagonal
- interface -> that is transport and framework layer
- service -> it is domain model that contains all relation business models and logic
without knowing about of source of or framework and transportation layer
- repository -> that contains source of data we have
*/
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
