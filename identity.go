package manager

// IdentityProvider specifies an API each concrete identity provider must
// fullfill (e.g. x509 certificates, JWT).
type IdentityProvider interface {
	// Key generates the platform access token.
	Key(string) (string, error)

	// IsValid determines whether or not provided token is valid.
	IsValid(string) bool
}
