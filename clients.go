package manager

import "strings"

// Client represents a Mainflux client. Each client is owned by one user, and
// it is assigned with the unique identifier and (temporary) access key.
type Client struct {
	Owner string            `json:"-"`
	ID    string            `json:"id"`
	Type  string            `json:"type"`
	Name  string            `json:"name,omitempty"`
	Key   string            `json:"key"`
	Meta  map[string]string `json:"meta,omitempty"`
}

var clientTypes map[string]bool = map[string]bool{
	"app":    true,
	"device": true,
}

func (c *Client) validate() error {
	if c.Type = strings.ToLower(c.Type); !clientTypes[c.Type] {
		return ErrMalformedClient
	}

	return nil
}

// ClientRepository specifies a client persistence API.
type ClientRepository interface {
	// Save persists the client. Successful operation is indicated by unique
	// identifier accompanied by nil error response. A non-nil error is
	// returned to indicate operation failure.
	Save(Client) (string, error)

	// Update performs an update to the existing client. A non-nil error is
	// returned to indicate operation failure.
	Update(Client) error

	// One retrieves the client having the provided identifier, that is owned
	// by the specified user.
	One(string, string) (Client, error)

	// All retrieves the clients owned by the specified user.
	All(string) []Client

	// Remove removes the client having the provided identifier, that is owned
	// by the specified user.
	Remove(string, string) error
}
