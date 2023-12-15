package greeter

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/iamsumit/example-go-kit/internal/service/greeter/store"
)

func Greet(g *store.Store) endpoint.Endpoint {
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
