package api

import (
	"fmt"
	"net/http"

	"github.com/mainflux/manager"
)

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
	return make(map[string]string)
}

func (tr tokenResponse) empty() bool {
	return tr.Token == ""
}

type createDeviceRequest struct {
	device manager.Device
	key    string
}

type createDeviceResponse struct {
	id uint
}

func (cdr createDeviceResponse) code() int {
	return http.StatusCreated
}

func (cdr createDeviceResponse) headers() map[string]string {
	return map[string]string{
		"Location": fmt.Sprintf("/devices/%d", cdr.id),
	}
}

func (cdr createDeviceResponse) empty() bool {
	return true
}
