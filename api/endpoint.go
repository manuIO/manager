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

func makeCreateDeviceEndpoint(s manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		cdr := request.(createDeviceRequest)

		id, err := s.CreateDevice(cdr.key, cdr.device)
		if err != nil {
			return nil, err
		}

		return createDeviceResponse{id}, nil
	}
}

func makeDeviceInfoEndpoint(s manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		dir := request.(deviceInfoRequest)

		device, err := s.DeviceInfo(dir.key, dir.id)
		if err != nil {
			return nil, err
		}

		return deviceInfoResponse{device}, nil
	}
}

func makeRemoveDeviceEndpoint(s manager.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		dir := request.(deviceInfoRequest)

		if err := s.RemoveDevice(dir.key, dir.id); err != nil {
			return nil, err
		}

		return deviceRemovalResponse{}, nil
	}
}
