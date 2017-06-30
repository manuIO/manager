package manager

// IdentityProvider specifies an API for identity management via security
// tokens.
type IdentityProvider interface {
	// Key generates the platform access token.
	Key(string) (string, error)

	// Identity extracts the entity identifier given its secret key.
	Identity(string) (string, error)
}
