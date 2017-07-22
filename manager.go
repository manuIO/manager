package manager

var _ Service = (*managerService)(nil)

type managerService struct {
	users   UserRepository
	clients ClientRepository
	hasher  Hasher
	idp     IdentityProvider
}

// NewService instantiates the domain service implementation.
func NewService(ur UserRepository, cr ClientRepository, hasher Hasher, idp IdentityProvider) Service {
	return &managerService{
		users:   ur,
		clients: cr,
		hasher:  hasher,
		idp:     idp,
	}
}

func (ms *managerService) Register(user User) error {
	if err := user.validate(); err != nil {
		return err
	}

	hash, err := ms.hasher.Hash(user.Password)
	if err != nil {
		return ErrInvalidCredentials
	}

	user.Password = hash
	return ms.users.Save(user)
}

func (ms *managerService) Login(user User) (string, error) {
	dbUser, err := ms.users.Get(user.Email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if err := ms.hasher.Compare(user.Password, dbUser.Password); err != nil {
		return "", ErrInvalidCredentials
	}

	return ms.idp.TemporaryKey(user.Email)
}

func (ms *managerService) CreateClient(key string, client Client) (string, error) {
	if err := client.validate(); err != nil {
		return "", err
	}

	sub, err := ms.idp.Identity(key)
	if err != nil {
		return "", err
	}

	client.Owner = sub
	client.Key, _ = ms.idp.PermanentKey(sub)

	return ms.clients.Save(client)
}

func (ms *managerService) ClientInfo(key string, id string) (Client, error) {
	sub, err := ms.idp.Identity(key)
	if err != nil {
		return Client{}, err
	}

	return ms.clients.One(sub, id)
}

func (ms *managerService) RemoveClient(key string, id string) error {
	sub, err := ms.idp.Identity(key)
	if err != nil {
		return err
	}

	return ms.clients.Remove(sub, id)
}
