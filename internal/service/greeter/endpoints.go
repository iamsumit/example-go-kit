package greeter

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/iamsumit/example-go-kit/internal/service/greeter/store"
)

type Endpoints struct {
	Greet endpoint.Endpoint
}

func MakeEndpoints(g *store.Store) Endpoints {
	return Endpoints{
		Greet: GreetEndpoint(g),
	}
}

func GreetEndpoint(g *store.Store) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		greet := request.(*ReqGreet)
		msg := g.Greet(greet.Message)
		if msg == "error" {
			return nil, errors.New("error while greeting")
		}
		return ResGreet{
			Message: msg,
		}, nil
	}
}
