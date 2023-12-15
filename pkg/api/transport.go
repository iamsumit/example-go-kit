package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func ServerOptions() []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(
			func(ctx context.Context, err error, w http.ResponseWriter) {
				EncodeResponse(ctx, w, err)
			},
		),
	}
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
	json.NewEncoder(w).Encode(
		ErrorResponse{
			Error: err.Error(),
		},
	)
}
