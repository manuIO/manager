package api

import (
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/mainflux/manager"
)

var _ manager.Service = (*metricService)(nil)

type metricService struct {
	counter metrics.Counter
	latency metrics.Histogram
	manager.Service
}

// NewMetricService instruments core service by tracking request count and
// latency.
func NewMetricService(counter metrics.Counter, latency metrics.Histogram, s manager.Service) manager.Service {
	return &metricService{
		counter: counter,
		latency: latency,
		Service: s,
	}
}

func (ms *metricService) Register(user manager.User) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "register").Add(1)
		ms.latency.With("method", "register").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.Service.Register(user)
}

func (ms *metricService) Login(user manager.User) (string, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "login").Add(1)
		ms.latency.With("method", "login").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.Service.Login(user)
}

func (ms *metricService) CreateDevice(key string, device manager.Device) (uint, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "create_device").Add(1)
		ms.latency.With("method", "create_device").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.Service.CreateDevice(key, device)
}

func (ms *metricService) DeviceInfo(key string, id uint) (manager.Device, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "device_info").Add(1)
		ms.latency.With("method", "device_info").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.Service.DeviceInfo(key, id)
}

func (ms *metricService) RemoveDevice(key string, id uint) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "remove_device").Add(1)
		ms.latency.With("method", "remove_device").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.Service.RemoveDevice(key, id)
}
