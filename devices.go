package manager

// Device represents a Mainflux device. Each device is owned by one user, and
// it is assigned with the unique identifier.
type Device struct {
	ID          uint
	Owner       string
	Name        string
	Description string
	Channels    []uint
}

// DeviceRepository specifies a device persistence API.
type DeviceRepository interface {
	// Save persists the device. Successful operation is indicated by unique
	// identifier accompanied by nil error response. A non-nil error is
	// returned to indicate operation failure.
	Save(Device) (uint, error)
}
