package mocks

import (
	"sync"

	"github.com/mainflux/manager"
)

var _ manager.ChannelRepository = (*channelRepositoryMock)(nil)

type channelRepositoryMock struct {
	mu       sync.Mutex
	counter  int
	channels map[string]manager.Channel
}

func NewChannelRepository() manager.ChannelRepository {
	return &channelRepositoryMock{
		channels: make(map[string]manager.Channel),
	}
}

func (cr *channelRepositoryMock) Save(channel manager.Channel) (string, error) {
	return "", nil
}

func (cr *channelRepositoryMock) Update(channel manager.Channel) error {
	return nil
}

func (cr *channelRepositoryMock) One(owner, id string) (manager.Channel, error) {
	return manager.Channel{}, nil
}

func (cr *channelRepositoryMock) All(owner string) []manager.Channel {
	return make([]manager.Channel, 0)
}

func (cr *channelRepositoryMock) Remove(owner, id string) error {
	return nil
}

func (cr *channelRepositoryMock) HasClient(owned, id, client string) bool {
	return false
}
