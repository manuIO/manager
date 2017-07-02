package mocks

import (
	"fmt"
	"sync"

	"github.com/mainflux/manager"
)

var _ manager.DeviceRepository = (*deviceRepositoryMock)(nil)

type deviceRepositoryMock struct {
	mu      sync.Mutex
	counter uint
	devices map[string]manager.Device
}

// NewDeviceRepository creates in-memory device repository.
func NewDeviceRepository() manager.DeviceRepository {
	return &deviceRepositoryMock{
		devices: make(map[string]manager.Device),
	}
}

func (dr *deviceRepositoryMock) Save(device manager.Device) (uint, error) {
	dr.mu.Lock()
	defer dr.mu.Unlock()

	if _, ok := dr.devices[key(device)]; ok {
		return dr.counter, nil
	}

	dr.counter += 1
	device.ID = dr.counter

	dr.devices[key(device)] = device

	return dr.counter, nil
}

func (dr *deviceRepositoryMock) One(owner string, id uint) (manager.Device, error) {
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
	return fmt.Sprintf("%d-%d", device.Owner, device.ID)
}

func (dr *deviceRepositoryMock) Remove(owner string, id uint) error {
	device := manager.Device{
		ID:    id,
		Owner: owner,
	}

	delete(dr.devices, key(device))

	return nil
}
