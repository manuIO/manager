package manager

var _ Service = (*managerService)(nil)

type managerService struct {
	users    UserRepository
	clients  ClientRepository
	channels ChannelRepository
	hasher   Hasher
	idp      IdentityProvider
}

// NewService instantiates the domain service implementation.
func NewService(users UserRepository, clients ClientRepository, channels ChannelRepository,
	hasher Hasher, idp IdentityProvider) Service {
	return &managerService{
		users:    users,
		clients:  clients,
		channels: channels,
		hasher:   hasher,
		idp:      idp,
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
	dbUser, err := ms.users.One(user.Email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if err := ms.hasher.Compare(user.Password, dbUser.Password); err != nil {
		return "", ErrInvalidCredentials
	}

	return ms.idp.TemporaryKey(user.Email)
}

func (ms *managerService) AddClient(key string, client Client) (string, error) {
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

func (ms *managerService) UpdateClient(key string, client Client) error {
	if err := client.validate(); err != nil {
		return err
	}

	sub, err := ms.idp.Identity(key)
	if err != nil {
		return err
	}

	client.Owner = sub

	return ms.clients.Update(client)
}

func (ms *managerService) ViewClient(key, id string) (Client, error) {
	sub, err := ms.idp.Identity(key)
	if err != nil {
		return Client{}, err
	}

	return ms.clients.One(sub, id)
}

func (ms *managerService) ListClients(key string) ([]Client, error) {
	sub, err := ms.idp.Identity(key)
	if err != nil {
		return nil, err
	}

	return ms.clients.All(sub), nil
}

func (ms *managerService) RemoveClient(key, id string) error {
	sub, err := ms.idp.Identity(key)
	if err != nil {
		return err
	}

	return ms.clients.Remove(sub, id)
}

func (ms *managerService) CreateChannel(key string, channel Channel) (string, error) {
	sub, err := ms.idp.Identity(key)
	if err != nil {
		return "", err
	}

	channel.Owner = sub
	return ms.channels.Save(channel)
}

func (ms *managerService) UpdateChannel(key string, channel Channel) error {
	return nil
}

func (ms *managerService) ViewChannel(key, id string) (Channel, error) {
	return Channel{}, nil
}

func (ms *managerService) ListChannels(key string) ([]Channel, error) {
	return make([]Channel, 0), nil
}

func (ms *managerService) RemoveChannel(key, id string) error {
	return nil
}

func (ms *managerService) CanRead(key, id string) bool {
	return false
}

func (ms *managerService) CanWrite(key, id string) bool {
	return false
}
