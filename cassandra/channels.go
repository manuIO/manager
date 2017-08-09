package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/mainflux/manager"
)

var _ manager.ChannelRepository = (*channelRepository)(nil)

type channelRepository struct {
	session *gocql.Session
}

// NewChannelRepository instantiates Cassandra channel repository.
func NewChannelRepository(session *gocql.Session) manager.ChannelRepository {
	return &channelRepository{session}
}

func (repo *channelRepository) Save(channel manager.Channel) (string, error) {
	cql := `INSERT INTO channels_by_user (user, id, name, connected) VALUES (?, ?, ?, ?)`
	id := gocql.TimeUUID()

	if err := repo.session.Query(cql, channel.Owner, id,
		channel.Name, channel.Connected).Exec(); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (repo *channelRepository) Update(channel manager.Channel) error {
	return nil
}

func (repo *channelRepository) One(owner, id string) (manager.Channel, error) {
	return manager.Channel{}, nil
}

func (repo *channelRepository) All(owner string) []manager.Channel {
	return make([]manager.Channel, 0)
}

func (repo *channelRepository) Remove(owner, id string) error {
	return nil
}

func (repo *channelRepository) HasClient(owned, id, client string) bool {
	return false
}
