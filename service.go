package manager

// Service specifies an API that must be fullfiled by the domain service
// implementation, and all of its decorators (e.g. logging & metrics).
type Service interface {
	// Register created new user account. In case the registration fails, a
	// non-nil error value is returned.
	Register(User) error

	// Login performs user authentication given its credentials. Successful
	// authentication generates new access token. Failed invocations are
	// identified by the non-nil error values present in the response.
	Login(User) (string, error)
}
