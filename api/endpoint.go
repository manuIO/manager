package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mainflux/manager"
)

func makeRegistrationEndpoint(s manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		user := request.(manager.User)
		err := s.Register(user)
		return tokenResponse{}, err
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

func makeCreateDeviceEndpoint(s manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		cdr := request.(createDeviceRequest)

		id, err := s.CreateDevice(cdr.key, cdr.device)
		if err != nil {
			return "", err
		}

		return createDeviceResponse{id}, nil
	}
}
