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
			return nil, err
		}

		return tokenResponse{token}, nil
	}
}

func makeCreateClientEndpoint(s manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		cdr := request.(createClientRequest)

		id, err := s.CreateClient(cdr.key, cdr.client)
		if err != nil {
			return nil, err
		}

		return createClientResponse{id}, nil
	}
}

func makeClientInfoEndpoint(s manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		cir := request.(clientInfoRequest)

		client, err := s.ClientInfo(cir.key, cir.id)
		if err != nil {
			return nil, err
		}

		return clientInfoResponse{client}, nil
	}
}

func makeRemoveClientEndpoint(s manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		cir := request.(clientInfoRequest)

		if err := s.RemoveClient(cir.key, cir.id); err != nil {
			return nil, err
		}

		return clientRemovalResponse{}, nil
	}
}
