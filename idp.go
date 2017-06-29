package manager

// IdentityProvider specifies an API each concrete identity provider must
// fulfill (e.g. x509 certificates, JWT).
type IdentityProvider interface {
	// Key generates the platform access token.
	Key(string) (string, error)

	// Identity extracts the entity identifier given its secret key.
	Identity(string) (string, error)
}
