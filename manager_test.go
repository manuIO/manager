package manager_test

import (
	"testing"

	"github.com/mainflux/manager"
	"github.com/mainflux/manager/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	ur  manager.UserRepository   = mocks.NewUserRepositoryMock()
	idp manager.IdentityProvider = mocks.NewIdentityProviderMock()
	svc manager.Service          = manager.NewService(ur, idp)
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
	}

	for _, tt := range cases {
		k, e := svc.Login(tt.user)
		assert.Equal(t, tt.key, k, "unexpected key retrieved")
		assert.Equal(t, tt.err, e, "unexpected error occurred")
	}
}
