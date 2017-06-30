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

func (u *User) validate() error {
	if u.Email == "" || u.Password == "" {
		return ErrInvalidCredentials
	}

	return nil
}

// UserRepository specifies an account persistence API.
type UserRepository interface {
	// Save persists the user account. A non-nil error is returned to indicate
	// operation failure.
	Save(User) error

	// Get retrieves user by its unique identifier (i.e. email).
	Get(string) (User, error)
}
