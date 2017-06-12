package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mainflux/manager"
)

type tokenResponse struct {
	Token string `json:"token"`
}

func makeRegistrationEndpoint(s manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		user := request.(manager.User)
		err := s.Register(user)
		return "", err
	}
}

func makeLoginEndpoint(s manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		user := request.(manager.User)

		token, err := s.Login(user)
		if err != nil {
			return "", err
		}

		return tokenResponse{token}, nil
	}
}
