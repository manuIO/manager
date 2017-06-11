package manager

// Service specifies an API that must be fullfiled by the domain service
// implementation, and all of its decorators (e.g. logging & metrics).
type Service interface {
	Register(User) error
	Login(string, string) (string, error)
}

// NewService instantiates the domain service implementation.
func NewService(users UserRepository, idp IdentityProvider) Service {
	return &managerService{
		users: users,
		idp:   idp,
	}
}
