package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	greetrepo "github.com/iamsumit/example-go-kit/internal/repository/greeter"
	"github.com/iamsumit/example-go-kit/internal/service/greeter"
	"github.com/iamsumit/example-go-kit/pkg/api"
)

type ErrorResponse struct {
	Errors struct {
		Error string `json:"error"`
	}
}

func main() {
	// greet repository
	greetRepo := greetrepo.New()
	greetService := greeter.NewService(greetRepo)
	r := mux.NewRouter()
	ge := greeter.MakeEndpoints(greetService)

	r.Methods("GET").Path("/greet/{message}").Handler(httptransport.NewServer(
		ge.Greet,
		greeter.Decode,
		api.EncodeResponse,
		api.ServerOptions()...,
	))

	// Start the HTTP server in a separate goroutine
	go func() {
		log.Println("Starting server on :4000...")
		if err := http.ListenAndServe(":4000", r); err != nil {
			log.Fatal(err)
		}
	}()

	// Listen for termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down gracefully...")
}
