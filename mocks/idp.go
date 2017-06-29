package mocks

import "github.com/mainflux/manager"

var _ manager.IdentityProvider = (*identityProviderMock)(nil)

type identityProviderMock struct{}

// NewIdentityProvider creates "mirror" identity provider, i.e. generated
// token will hold value provided by the caller.
func NewIdentityProvider() manager.IdentityProvider {
	return &identityProviderMock{}
}

func (idp *identityProviderMock) Key(id string) (string, error) {
	return id, nil
}

func (idp *identityProviderMock) Identity(key string) (string, error) {
	return key, nil
}
