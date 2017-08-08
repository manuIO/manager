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
		return tokenRes{}, err
	}
}

func loginEndpoint(svc manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		user := request.(manager.User)

		token, err := svc.Login(user)
		if err != nil {
			return nil, err
		}

		return tokenRes{token}, nil
	}
}

func addClientEndpoint(svc manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addClientReq)

		id, err := svc.AddClient(req.key, req.client)
		if err != nil {
			return nil, err
		}

		return addClientRes{id}, nil
	}
}

func viewClientEndpoint(svc manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(viewResourceReq)

		client, err := svc.ViewClient(req.key, req.id)
		if err != nil {
			return nil, err
		}

		return viewClientRes{client}, nil
	}
}

func listClientsEndpoint(svc manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(listResourcesReq)

		clients, err := svc.ListClients(req.key)
		if err != nil {
			return nil, err
		}

		return listClientsRes{clients, len(clients)}, nil
	}
}

func removeClientEndpoint(svc manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(viewResourceReq)

		if err := svc.RemoveClient(req.key, req.id); err != nil {
			return nil, err
		}

		return removeClientRes{}, nil
	}
}
