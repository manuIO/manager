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

type addClientReq struct {
	key    string
	client manager.Client
}

type viewResourceReq struct {
	key string
	id  string
}

type listResourcesReq struct {
	key    string
	size   int
	offset int
}

type tokenRes struct {
	Token string `json:"token,omitempty"`
}

func (res tokenRes) code() int {
	return http.StatusCreated
}

func (res tokenRes) headers() map[string]string {
	return map[string]string{}
}

func (res tokenRes) empty() bool {
	return res.Token == ""
}

type addClientRes struct {
	id string
}

func (res addClientRes) code() int {
	return http.StatusCreated
}

func (res addClientRes) headers() map[string]string {
	return map[string]string{
		"Location": fmt.Sprint("/clients/", res.id),
	}
}

func (res addClientRes) empty() bool {
	return true
}

type viewClientRes struct {
	manager.Client
}

func (res viewClientRes) code() int {
	return http.StatusOK
}

func (res viewClientRes) headers() map[string]string {
	return map[string]string{}
}

func (res viewClientRes) empty() bool {
	return false
}

type listClientsRes struct {
	Clients []manager.Client `json:"clients"`
	count   int
}

func (res listClientsRes) code() int {
	return http.StatusOK
}

func (res listClientsRes) headers() map[string]string {
	return map[string]string{
		"X-Count": fmt.Sprintf("%d", res.count),
	}
}

func (res listClientsRes) empty() bool {
	return false
}

type removeClientRes struct{}

func (res removeClientRes) code() int {
	return http.StatusNoContent
}

func (res removeClientRes) headers() map[string]string {
	return map[string]string{}
}

func (res removeClientRes) empty() bool {
	return true
}
