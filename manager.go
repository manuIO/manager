package manager

var _ Service = (*managerService)(nil)

type managerService struct {
	users UserRepository
	idp   IdentityProvider
}

// Register created new user account. In case the registration fails, a non-nil
// error value is returned.
func (ms *managerService) Register(user User) error {
	return ms.users.Save(user)
}

// Login performs user authentication given its credentials. Successful
// authentication generates new access token. Failed invocations are identified
// by the non-nil error values present in the response.
func (ms *managerService) Login(email, password string) (string, error) {
	if e := ms.users.Exists(email, password); !e {
		return "", ErrInvalidCredentials
	}

	return ms.idp.Key(email)
}
