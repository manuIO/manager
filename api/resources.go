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

type deviceInfoRequest struct {
	id  uint
	key string
}

type deviceInfoResponse struct {
	manager.Device
}

func (dir deviceInfoResponse) code() int {
	return http.StatusOK
}

func (dir deviceInfoResponse) headers() map[string]string {
	return map[string]string{}
}

func (dir deviceInfoResponse) empty() bool {
	return false
}

type deviceRemovalResponse struct {
}

func (drr deviceRemovalResponse) code() int {
	return http.StatusNoContent
}

func (drr deviceRemovalResponse) headers() map[string]string {
	return map[string]string{}
}

func (drr deviceRemovalResponse) empty() bool {
	return true
}
