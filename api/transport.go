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

	clientCreationHandler := kithttp.NewServer(
		makeCreateClientEndpoint(svc),
		decodeCreateClientRequest,
		encodeResponse,
		opts...,
	)

	clientInfoHandler := kithttp.NewServer(
		makeClientInfoEndpoint(svc),
		decodeClientInfoRequest,
		encodeResponse,
		opts...,
	)

	removeClientHandler := kithttp.NewServer(
		makeRemoveClientEndpoint(svc),
		decodeClientInfoRequest,
		encodeResponse,
		opts...,
	)

	r := bone.New()

	r.Post("/users", registrationHandler)
	r.Post("/tokens", loginHandler)
	r.Post("/clients", clientCreationHandler)
	r.Get("/clients/:id", clientInfoHandler)
	r.Delete("/clients/:id", removeClientHandler)
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

func decodeCreateClientRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var client manager.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		return nil, err
	}

	cdr := createClientRequest{
		key:    r.Header.Get("Authorization"),
		client: client,
	}

	return cdr, nil
}

func decodeClientInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	cir := clientInfoRequest{
		key: r.Header.Get("Authorization"),
		id:  bone.GetValue(r, "id"),
	}

	return cir, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", contentType)

	if ar, ok := response.(apiResponse); ok {
		for k, v := range ar.headers() {
			w.Header().Set(k, v)
		}

		w.WriteHeader(ar.code())

		if ar.empty() {
			return nil
		}
	}

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", contentType)

	switch err {
	case manager.ErrInvalidCredentials, manager.ErrMalformedClient:
		w.WriteHeader(http.StatusBadRequest)
	case manager.ErrUnauthorizedAccess:
		w.WriteHeader(http.StatusForbidden)
	case manager.ErrNotFound:
		w.WriteHeader(http.StatusNotFound)
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
