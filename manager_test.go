package manager_test

import (
	"testing"

	"github.com/mainflux/manager"
	"github.com/mainflux/manager/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	users   manager.UserRepository   = mocks.NewUserRepository()
	clients manager.ClientRepository = mocks.NewClientRepository()
	hasher  manager.Hasher           = mocks.NewHasher()
	idp     manager.IdentityProvider = mocks.NewIdentityProvider()
	svc     manager.Service          = manager.NewService(users, clients, hasher, idp)
)

func TestRegister(t *testing.T) {
	var cases = []struct {
		user manager.User
		err  error
	}{
		{manager.User{"foo@bar.com", "pass"}, nil},
		{manager.User{"foo@bar.com", "pass"}, manager.ErrConflict},
		{manager.User{"", "pass"}, manager.ErrInvalidCredentials},
		{manager.User{"abc@bar.com", ""}, manager.ErrInvalidCredentials},
		{manager.User{"abc@bar.com", "pass"}, nil},
	}

	for _, tc := range cases {
		e := svc.Register(tc.user)
		assert.Equal(t, tc.err, e, "unexpected error occurred")
	}
}

func TestLogin(t *testing.T) {
	var cases = []struct {
		user manager.User
		key  string
		err  error
	}{
		{manager.User{"foo@bar.com", "pass"}, "foo@bar.com", nil},
		{manager.User{"new@bar.com", "pass"}, "", manager.ErrInvalidCredentials},
		{manager.User{"foo@bar.com", ""}, "", manager.ErrInvalidCredentials},
	}

	for _, tc := range cases {
		k, e := svc.Login(tc.user)
		assert.Equal(t, tc.key, k, "unexpected key retrieved")
		assert.Equal(t, tc.err, e, "unexpected error occurred")
	}
}

func TestAddClient(t *testing.T) {
	var cases = []struct {
		key    string
		client manager.Client
		id     string
		err    error
	}{
		{"foo@bar.com", manager.Client{Name: "a"}, "1", nil},
		{"foo@bar.com", manager.Client{Name: "b"}, "2", nil},
		{"foo@bar.com", manager.Client{Name: "c"}, "3", nil},
		{"foo@bar.com", manager.Client{ID: "3", Name: "c"}, "3", nil},
		{"", manager.Client{Name: "d"}, "", manager.ErrUnauthorizedAccess},
		{"foo@bar.com", manager.Client{}, "", manager.ErrMalformedClient},
	}

	for _, tc := range cases {
		id, err := svc.AddClient(tc.key, tc.client)
		assert.Equal(t, tc.id, id, "unexpected id retrieved")
		assert.Equal(t, tc.err, err, "unexpected error occurred")
	}
}

func TestViewClient(t *testing.T) {
	var cases = []struct {
		id  string
		key string
		err error
	}{
		{"1", "foo@bar.com", nil},
		{"1", "", manager.ErrUnauthorizedAccess},
		{"5", "foo@bar.com", manager.ErrNotFound},
	}

	for _, tc := range cases {
		_, err := svc.ViewClient(tc.key, tc.id)
		assert.Equal(t, tc.err, err, "unexpected error occurred")
	}
}

func TestRemoveClient(t *testing.T) {
	var cases = []struct {
		id  string
		key string
		err error
	}{
		{"1", "", manager.ErrUnauthorizedAccess},
		{"1", "foo@bar.com", nil},
		{"1", "foo@bar.com", nil},
		{"2", "foo@bar.com", nil},
	}

	for _, tc := range cases {
		err := svc.RemoveClient(tc.key, tc.id)
		assert.Equal(t, tc.err, err, "unexpected error occurred")
	}
}
