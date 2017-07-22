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

func (ms *metricService) CreateClient(key string, client manager.Client) (string, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "create_client").Add(1)
		ms.latency.With("method", "create_client").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.Service.CreateClient(key, client)
}

func (ms *metricService) ClientInfo(key string, id string) (manager.Client, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "client_info").Add(1)
		ms.latency.With("method", "client_info").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.Service.ClientInfo(key, id)
}

func (ms *metricService) RemoveClient(key string, id string) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "remove_client").Add(1)
		ms.latency.With("method", "remove_client").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.Service.RemoveClient(key, id)
}
