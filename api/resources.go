package api

import (
	"fmt"
	"net/http"

	"github.com/mainflux/manager"
)

const contentType = "application/json; charset=utf-8"

type apiResponse interface {
	code() int
	headers() map[string]string
	empty() bool
}

type tokenResponse struct {
	Token string `json:"token,omitempty"`
}

func (tr tokenResponse) code() int {
	return http.StatusCreated
}

func (tr tokenResponse) headers() map[string]string {
	return map[string]string{}
}

func (tr tokenResponse) empty() bool {
	return tr.Token == ""
}

type createClientRequest struct {
	client manager.Client
	key    string
}

type createClientResponse struct {
	id string
}

func (ccr createClientResponse) code() int {
	return http.StatusCreated
}

func (ccr createClientResponse) headers() map[string]string {
	return map[string]string{
		"Location": fmt.Sprint("/clients/", ccr.id),
	}
}

func (ccr createClientResponse) empty() bool {
	return true
}

type clientInfoRequest struct {
	id  string
	key string
}

type clientInfoResponse struct {
	manager.Client
}

func (cir clientInfoResponse) code() int {
	return http.StatusOK
}

func (cir clientInfoResponse) headers() map[string]string {
	return map[string]string{}
}

func (cir clientInfoResponse) empty() bool {
	return false
}

type clientRemovalResponse struct {
}

func (crr clientRemovalResponse) code() int {
	return http.StatusNoContent
}

func (crr clientRemovalResponse) headers() map[string]string {
	return map[string]string{}
}

func (crr clientRemovalResponse) empty() bool {
	return true
}
