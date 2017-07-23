package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mainflux/manager"
)

func registrationEndpoint(svc manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		user := request.(manager.User)
		err := svc.Register(user)
		return tokenRep{}, err
	}
}

func loginEndpoint(svc manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		user := request.(manager.User)

		token, err := svc.Login(user)
		if err != nil {
			return nil, err
		}

		return tokenRep{token}, nil
	}
}

func addClientEndpoint(svc manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addClientReq)

		id, err := svc.AddClient(req.key, req.client)
		if err != nil {
			return nil, err
		}

		return addClientRep{id}, nil
	}
}

func viewClientEndpoint(svc manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(viewClientReq)

		client, err := svc.ViewClient(req.key, req.id)
		if err != nil {
			return nil, err
		}

		return viewClientRep{client}, nil
	}
}

func removeClientEndpoint(svc manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(viewClientReq)

		if err := svc.RemoveClient(req.key, req.id); err != nil {
			return nil, err
		}

		return removeClientRep{}, nil
	}
}
