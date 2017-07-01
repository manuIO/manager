package manager_test

import (
	"testing"

	"github.com/mainflux/manager"
	"github.com/mainflux/manager/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	users   manager.UserRepository   = mocks.NewUserRepository()
	devices manager.DeviceRepository = mocks.NewDeviceRepository()
	hasher  manager.Hasher           = mocks.NewHasher()
	idp     manager.IdentityProvider = mocks.NewIdentityProvider()
	svc     manager.Service          = manager.NewService(users, devices, hasher, idp)
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

func TestCreateDevice(t *testing.T) {
	var cases = []struct {
		key    string
		device manager.Device
		id     uint
		err    error
	}{
		{"foo@bar.com", manager.Device{Name: "a"}, 1, nil},
		{"foo@bar.com", manager.Device{Name: "b"}, 2, nil},
		{"foo@bar.com", manager.Device{Name: "c"}, 3, nil},
		{"foo@bar.com", manager.Device{ID: 3, Name: "c"}, 3, nil},
		{"", manager.Device{Name: "d"}, 0, manager.ErrUnauthorizedAccess},
		{"foo@bar.com", manager.Device{}, 0, manager.ErrMalformedDevice},
	}

	for _, tc := range cases {
		id, err := svc.CreateDevice(tc.key, tc.device)
		assert.Equal(t, tc.id, id, "unexpected id retrieved")
		assert.Equal(t, tc.err, err, "unexpected error occurred")
	}
}

func TestDeviceInfo(t *testing.T) {
	var cases = []struct {
		id  uint
		key string
		err error
	}{
		{1, "foo@bar.com", nil},
		{1, "", manager.ErrUnauthorizedAccess},
		{5, "foo@bar.com", manager.ErrNotFound},
	}

	for _, tc := range cases {
		_, err := svc.DeviceInfo(tc.key, tc.id)
		assert.Equal(t, tc.err, err, "unexpected error occurred")
	}
}
