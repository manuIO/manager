package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/mainflux/manager"
)

var _ manager.ClientRepository = (*clientRepository)(nil)

type clientRepository struct {
	session gocql.Session
}

// NewClientRepository instantiates Cassandra client repository.
func NewClientRepository(session gocql.Session) manager.ClientRepository {
	return &clientRepository{session}
}

func (dr *clientRepository) Save(client manager.Client) (string, error) {
	return "", manager.ErrConflict
}

func (dr *clientRepository) One(owner string, id string) (manager.Client, error) {
	return manager.Client{}, manager.ErrConflict
}

func (dr *clientRepository) Remove(owner string, id string) error {
	return nil
}
