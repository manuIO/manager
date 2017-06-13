package mocks

import "github.com/mainflux/manager"

var _ manager.IdentityProvider = (*identityProviderMock)(nil)

type identityProviderMock struct{}

// NewIdentityProvider creates test-friendly identity provider.
func NewIdentityProvider() manager.IdentityProvider {
	return &identityProviderMock{}
}

func (idp *identityProviderMock) Key(id string) (string, error) {
	return id, nil
}

func (idp *identityProviderMock) IsValid(key string) bool {
	return len(key) > 0
}
