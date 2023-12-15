package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/iamsumit/example-go-kit/internal/service/greeter"
	greetstore "github.com/iamsumit/example-go-kit/internal/service/greeter/store"
)

type ErrorResponse struct {
	Errors struct {
		Error string `json:"error"`
	}
}

func main() {
	// greet store
	g := greetstore.New()

	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(
			func(ctx context.Context, err error, w http.ResponseWriter) {
				kitlog.NewLogfmtLogger(os.Stderr).Log("error", err.Error())
				EncodeResponse(ctx, w, err)
			},
		),
	}

	r.Methods("GET").Path("/greet/{message}").Handler(httptransport.NewServer(
		greeter.Greet(g),
		greeter.Decode,
		EncodeResponse,
		options...,
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

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Printf("response: %v", response)
	if e, ok := response.(error); ok {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e, w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
