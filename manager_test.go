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

	for _, tt := range cases {
		e := svc.Register(tt.user)
		assert.Equal(t, tt.err, e, "unexpected error occurred")
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

	for _, tt := range cases {
		k, e := svc.Login(tt.user)
		assert.Equal(t, tt.key, k, "unexpected key retrieved")
		assert.Equal(t, tt.err, e, "unexpected error occurred")
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
		{"", manager.Device{Name: "d"}, 0, manager.ErrInvalidCredentials},
		{"foo@bar.com", manager.Device{}, 0, manager.ErrMalformedDevice},
	}

	for _, tt := range cases {
		id, err := svc.CreateDevice(tt.key, tt.device)
		assert.Equal(t, tt.id, id, "unexpected id retrieved")
		assert.Equal(t, tt.err, err, "unexpected error occurred")
	}
}
