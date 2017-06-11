package manager

import "errors"

var (
	// ErrConflict indicates usage of the existing email during account
	// registration.
	ErrConflict error = errors.New("email already taken")

	// ErrInvalidCredentials indicates malformed account credentials.
	ErrInvalidCredentials error = errors.New("invalid email or password")
)

// User represents a Mainflux user account. Each user is identified given its
// email and password.
type User struct {
	Email    string
	Password string
}

// UserRepository specifies an API that needs to be implemented by the concrete
// storage providers (e.g. CockroachDB repository).
type UserRepository interface {
	// Save persists the user account. A non-nil error is returned to indicate
	// operation failure.
	Save(User) error

	// Exists determines whether or not an account with given credentials
	// exists in the system.
	Exists(string, string) bool
}
