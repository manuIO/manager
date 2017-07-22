package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/mainflux/manager"
)

var _ manager.UserRepository = (*userRepository)(nil)

type userRepository struct {
	session *gocql.Session
}

// NewUserRepository instantiates Cassandra user repository.
func NewUserRepository(session *gocql.Session) manager.UserRepository {
	return &userRepository{session}
}

func (ur *userRepository) Save(user manager.User) error {
	return manager.ErrConflict
}

func (ur *userRepository) Get(email string) (manager.User, error) {
	return manager.User{}, manager.ErrConflict
}
