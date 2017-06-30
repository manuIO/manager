package manager

var _ Service = (*managerService)(nil)

type managerService struct {
	users   UserRepository
	devices DeviceRepository
	hasher  Hasher
	idp     IdentityProvider
}

// NewService instantiates the domain service implementation.
func NewService(ur UserRepository, dr DeviceRepository, hasher Hasher, idp IdentityProvider) Service {
	return &managerService{
		users:   ur,
		devices: dr,
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

	return ms.idp.Key(user.Email)
}

func (ms *managerService) CreateDevice(key string, device Device) (uint, error) {
	if err := device.validate(); err != nil {
		return 0, err
	}

	sub, err := ms.idp.Identity(key)
	if err != nil {
		return 0, err
	}

	device.Owner = sub

	return ms.devices.Save(device)
}
