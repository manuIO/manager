package manager

import "errors"

var (
	// ErrConflict indicates usage of the existing email during account
	// registration.
	ErrConflict error = errors.New("email already taken")

	// ErrInvalidCredentials indicates malformed account credentials.
	ErrInvalidCredentials error = errors.New("invalid email or password")

	// ErrMalformedDevice indicates malformed device specification (e.g. empty name).
	ErrMalformedDevice error = errors.New("malformed device specification")

	// ErrUnauthorizedAccess indicates missing or invalid credentials provided
	// when accessing a protected resource.
	ErrUnauthorizedAccess error = errors.New("missing or invalid credentials provided")

	// ErrNotFound indicates a non-existent entity request.
	ErrNotFound error = errors.New("non-existent entity")
)

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

	// CreateDevice adds new device to the user identified by the provided key.
	CreateDevice(string, Device) (string, error)

	// DeviceInfo retrieves data about the device belonging to the user
	// identified by the provided key having the provided device ID.
	DeviceInfo(string, string) (Device, error)

	// RemoveDevice removes device belonging to the user identified by the
	// provided key having the provided device ID.
	RemoveDevice(string, string) error
}
