package cassandra

import (
	"log"

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
	cql := `INSERT INTO users (email, password) VALUES (?, ?) IF NOT EXISTS`

	applied, err := ur.session.Query(cql, user.Email, user.Password).ScanCAS()
	if !applied {
		return manager.ErrConflict
	}

	return err
}

func (ur *userRepository) One(email string) (manager.User, error) {
	cql := `SELECT email, password FROM users WHERE email = ?`

	user := manager.User{}
	row := map[string]interface{}{
		"email":    &user.Email,
		"password": &user.Password,
	}

	iter := ur.session.Query(cql, email).Iter()
	if !iter.MapScan(row) {
		return user, manager.ErrInvalidCredentials
	}

	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	return user, nil
}
