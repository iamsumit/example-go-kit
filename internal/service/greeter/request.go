package greeter

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type ReqGreet struct {
	Message string `json:"message"`
}

func Decode(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	msg, ok := vars["message"]
	if !ok {
		return nil, errors.New("message not found")
	}

	if msg == "nodecode" {
		return nil, errors.New("error while decoding")
	}

	return &ReqGreet{
		Message: msg,
	}, nil
}
