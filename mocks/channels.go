package mocks

import (
	"strconv"
	"sync"

	"github.com/mainflux/manager"
)

var _ manager.ChannelRepository = (*channelRepositoryMock)(nil)

type channelRepositoryMock struct {
	mu       sync.Mutex
	counter  int
	channels map[string]manager.Channel
}

// NewChannelRepository creates in-memory channel repository.
func NewChannelRepository() manager.ChannelRepository {
	return &channelRepositoryMock{
		channels: make(map[string]manager.Channel),
	}
}

func (repo *channelRepositoryMock) Save(channel manager.Channel) (string, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.counter += 1
	channel.ID = strconv.Itoa(repo.counter)

	repo.channels[key(channel.Owner, channel.ID)] = channel

	return channel.ID, nil
}

func (repo *channelRepositoryMock) Update(channel manager.Channel) error {
	return nil
}

func (repo *channelRepositoryMock) One(owner, id string) (manager.Channel, error) {
	return manager.Channel{}, nil
}

func (repo *channelRepositoryMock) All(owner string) []manager.Channel {
	return make([]manager.Channel, 0)
}

func (repo *channelRepositoryMock) Remove(owner, id string) error {
	return nil
}

func (repo *channelRepositoryMock) HasClient(owned, id, client string) bool {
	return false
}
