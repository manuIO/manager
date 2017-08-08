package manager

import "errors"

var (
	// ErrConflict indicates usage of the existing email during account
	// registration.
	ErrConflict error = errors.New("email already taken")

	// ErrInvalidCredentials indicates malformed account credentials.
	ErrInvalidCredentials error = errors.New("invalid email or password")

	// ErrMalformedClient indicates malformed client specification (e.g. empty name).
	ErrMalformedClient error = errors.New("malformed client specification")

	// ErrUnauthorizedAccess indicates missing or invalid credentials provided
	// when accessing a protected resource.
	ErrUnauthorizedAccess error = errors.New("missing or invalid credentials provided")

	// ErrNotFound indicates a non-existent entity request.
	ErrNotFound error = errors.New("non-existent entity")
)

// Service specifies an API that must be fullfiled by the domain service
// implementation, and all of its decorators (e.g. logging & metrics).
type Service interface {
	// Register creates new user account. In case of the failed registration, a
	// non-nil error value is returned.
	Register(User) error

	// Login authenticates the user given its credentials. Successful
	// authentication generates new access token. Failed invocations are
	// identified by the non-nil error values in the response.
	Login(User) (string, error)

	// AddClient adds new client to the user identified by the provided key.
	AddClient(string, Client) (string, error)

	// ViewClient retrieves data about the client identified with the provided
	// ID, that belongs to the user identified by the provided key.
	ViewClient(string, string) (Client, error)

	// ListClients retrieves data about all clients that belong to the user
	// identifier by the provided key.
	ListClients(string) ([]Client, error)

	// RemoveClient removes client identified with the provided ID, that
	// belongs to the user identified by the provided key.
	RemoveClient(string, string) error
}
