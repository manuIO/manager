package api

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-zoo/bone"
	"github.com/mainflux/manager"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MakeHandler returns a HTTP handler for API endpoints.
func MakeHandler(svc manager.Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	registrationHandler := kithttp.NewServer(
		makeRegistrationEndpoint(svc),
		decodeCredentialsRequest,
		encodeResponse,
		opts...,
	)

	loginHandler := kithttp.NewServer(
		makeLoginEndpoint(svc),
		decodeCredentialsRequest,
		encodeResponse,
		opts...,
	)

	deviceCreationHandler := kithttp.NewServer(
		makeCreateDeviceEndpoint(svc),
		decodeCreateDeviceRequest,
		encodeResponse,
		opts...,
	)

	r := bone.New()

	r.Post("/users", registrationHandler)
	r.Post("/tokens", loginHandler)
	r.Post("/devices", deviceCreationHandler)
	r.Handle("/metrics", promhttp.Handler())

	return r
}

func decodeCredentialsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var user manager.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func decodeCreateDeviceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var device manager.Device
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		return nil, err
	}

	cdr := createDeviceRequest{
		key:    r.Header.Get("Authorization"),
		device: device,
	}

	return cdr, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if ar, ok := response.(apiResponse); ok {
		w.WriteHeader(ar.code())

		for k, v := range ar.headers() {
			w.Header().Set(k, v)
		}

		if ar.empty() {
			return nil
		}
	}

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch err {
	case manager.ErrInvalidCredentials, manager.ErrMalformedDevice:
		w.WriteHeader(http.StatusBadRequest)
	case manager.ErrUnauthorizedAccess:
		w.WriteHeader(http.StatusForbidden)
	case manager.ErrConflict:
		w.WriteHeader(http.StatusConflict)
	default:
		if _, ok := err.(*json.SyntaxError); ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
	}
}
