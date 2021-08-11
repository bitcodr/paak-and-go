package transport

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"

	tripservice "github.com/bitcodr/paak-and-go/internal/domain/service/trip"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/config"
	tripcontroller "github.com/bitcodr/paak-and-go/internal/interfaces/controller/trip"
	"github.com/bitcodr/paak-and-go/internal/interfaces/middleware"
)

//Service we can add all our services in here and pass it to our transport patterns
type Service struct {
	TripService tripservice.ITrip

	//register your services in transport in here
}

//InitRest to initialise http apis
//it is possible in a project that use multiple transport like grpc, http, etc
func InitRest(ctx context.Context, services *Service, config *config.Service) {
	router := mux.NewRouter()

	router.HandleFunc("/trip", tripcontroller.List(ctx, services.TripService)).Methods(http.MethodGet)
	router.HandleFunc("/trip/{id}", tripcontroller.Show(ctx, services.TripService)).Methods(http.MethodGet)
	router.HandleFunc("/trip", tripcontroller.Store(ctx, services.TripService)).Methods(http.MethodPost)

	router.Use(middleware.Logging)

	srv := &http.Server{
		Addr:         config.Host + ":" + config.RestPort,
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
		IdleTimeout:  config.IdleTimeout,
		Handler:      router,
	}

	go func() {
		fmt.Printf("listening on %s\n", config.Host+":"+config.RestPort)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	graceFullyShotDown(ctx, srv, config)
}

//graceFullyShotDown terminate all process and http services with ctrl+c
func graceFullyShotDown(ctx context.Context, srv *http.Server, config *config.Service) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(ctx, config.IdleTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
		return
	}
	//todo call db connection close
	log.Println("shutting down")
	os.Exit(0)
}
