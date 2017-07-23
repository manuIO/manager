package mocks

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/mainflux/manager"
)

var _ manager.ClientRepository = (*clientRepositoryMock)(nil)

type clientRepositoryMock struct {
	mu      sync.Mutex
	counter int
	clients map[string]manager.Client
}

// NewClientRepository creates in-memory client repository.
func NewClientRepository() manager.ClientRepository {
	return &clientRepositoryMock{
		clients: make(map[string]manager.Client),
	}
}

func (cr *clientRepositoryMock) Save(client manager.Client) (string, error) {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	if c, ok := cr.clients[key(client)]; ok {
		return c.ID, nil
	}

	cr.counter += 1
	client.ID = strconv.Itoa(cr.counter)

	cr.clients[key(client)] = client

	return client.ID, nil
}

func (cr *clientRepositoryMock) One(owner string, id string) (manager.Client, error) {
	client := manager.Client{
		ID:    id,
		Owner: owner,
	}

	if c, ok := cr.clients[key(client)]; ok {
		return c, nil
	}

	return manager.Client{}, manager.ErrNotFound
}

func key(client manager.Client) string {
	return fmt.Sprintf("%s-%s", client.Owner, client.ID)
}

func (cr *clientRepositoryMock) Remove(owner string, id string) error {
	client := manager.Client{
		ID:    id,
		Owner: owner,
	}

	delete(cr.clients, key(client))

	return nil
}
