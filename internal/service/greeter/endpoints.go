package greeter

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Greet endpoint.Endpoint
}

func MakeEndpoints(g Service) Endpoints {
	return Endpoints{
		Greet: GreetEndpoint(g),
	}
}

func GreetEndpoint(g Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		greet := request.(*ReqGreet)
		msg, err := g.Greet(greet.Message)
		if err != nil {
			return nil, err
		}

		return ResGreet{
			Message: msg,
		}, nil
	}
}
