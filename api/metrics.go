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

func (s *metricService) Register(user manager.User) (err error) {
	defer func(begin time.Time) {
		s.counter.With("method", "register").Add(1)
		s.latency.With("method", "register").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Register(user)
}

func (s *metricService) Login(user manager.User) (token string, err error) {
	defer func(begin time.Time) {
		s.counter.With("method", "login").Add(1)
		s.latency.With("method", "login").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Login(user)
}
