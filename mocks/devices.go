package mocks

import (
	"sync"

	"github.com/mainflux/manager"
)

var _ manager.DeviceRepository = (*deviceRepositoryMock)(nil)

type deviceRepositoryMock struct {
	mu      sync.Mutex
	devices map[string][]manager.Device
}

func NewDeviceRepository() manager.DeviceRepository {
	return &deviceRepositoryMock{
		devices: make(map[string][]manager.Device),
	}
}

func (dr *deviceRepositoryMock) Save(device manager.Device) (uint, error) {
	dr.mu.Lock()
	defer dr.mu.Unlock()

	dr.devices[device.Owner] = append(dr.devices[device.Owner], device)
	return uint(len(dr.devices[device.Owner])), nil
}
