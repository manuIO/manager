package manager

// Device represents a Mainflux device. Each device is owned by one user, and
// it is assigned with the unique identifier.
type Device struct {
	ID          uint   `json:"id"`
	Owner       string `json:"-"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description"`
	Channels    []uint `json:"channels"`
}

func (d *Device) validate() error {
	if d.Name == "" || len(d.Name) > 50 {
		return ErrMalformedDevice
	}

	return nil
}

// DeviceRepository specifies a device persistence API.
type DeviceRepository interface {
	// Save persists the device. Successful operation is indicated by unique
	// identifier accompanied by nil error response. A non-nil error is
	// returned to indicate operation failure.
	Save(Device) (uint, error)

	// One retrieves the device identified by the provided unique ID and owned
	// by the specified user.
	One(uint, string) (Device, error)
}
