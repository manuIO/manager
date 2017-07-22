package mocks

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/mainflux/manager"
)

var _ manager.DeviceRepository = (*deviceRepositoryMock)(nil)

type deviceRepositoryMock struct {
	mu      sync.Mutex
	counter int
	devices map[string]manager.Device
}

// NewDeviceRepository creates in-memory device repository.
func NewDeviceRepository() manager.DeviceRepository {
	return &deviceRepositoryMock{
		devices: make(map[string]manager.Device),
	}
}

func (dr *deviceRepositoryMock) Save(device manager.Device) (string, error) {
	dr.mu.Lock()
	defer dr.mu.Unlock()

	if d, ok := dr.devices[key(device)]; ok {
		return d.ID, nil
	}

	dr.counter += 1
	device.ID = strconv.Itoa(dr.counter)

	dr.devices[key(device)] = device

	return device.ID, nil
}

func (dr *deviceRepositoryMock) One(owner string, id string) (manager.Device, error) {
	device := manager.Device{
		ID:    id,
		Owner: owner,
	}

	if device, ok := dr.devices[key(device)]; ok {
		return device, nil
	}

	return manager.Device{}, manager.ErrNotFound
}

func key(device manager.Device) string {
	return fmt.Sprintf("%s-%s", device.Owner, device.ID)
}

func (dr *deviceRepositoryMock) Remove(owner string, id string) error {
	device := manager.Device{
		ID:    id,
		Owner: owner,
	}

	delete(dr.devices, key(device))

	return nil
}
