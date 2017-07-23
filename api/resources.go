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

type tokenRep struct {
	Token string `json:"token,omitempty"`
}

func (rep tokenRep) code() int {
	return http.StatusCreated
}

func (rep tokenRep) headers() map[string]string {
	return map[string]string{}
}

func (rep tokenRep) empty() bool {
	return rep.Token == ""
}

type addClientReq struct {
	client manager.Client
	key    string
}

type addClientRep struct {
	id string
}

func (rep addClientRep) code() int {
	return http.StatusCreated
}

func (rep addClientRep) headers() map[string]string {
	return map[string]string{
		"Location": fmt.Sprint("/clients/", rep.id),
	}
}

func (rep addClientRep) empty() bool {
	return true
}

type viewClientReq struct {
	id  string
	key string
}

type viewClientRep struct {
	manager.Client
}

func (rep viewClientRep) code() int {
	return http.StatusOK
}

func (rep viewClientRep) headers() map[string]string {
	return map[string]string{}
}

func (rep viewClientRep) empty() bool {
	return false
}

type removeClientRep struct{}

func (rep removeClientRep) code() int {
	return http.StatusNoContent
}

func (rep removeClientRep) headers() map[string]string {
	return map[string]string{}
}

func (rep removeClientRep) empty() bool {
	return true
}
