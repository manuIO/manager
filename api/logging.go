package api

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/mainflux/manager"
)

var _ manager.Service = (*loggingService)(nil)

type loggingService struct {
	logger log.Logger
	manager.Service
}

// NewLoggingService adds logging facilities to the core service.
func NewLoggingService(logger log.Logger, s manager.Service) manager.Service {
	return &loggingService{logger, s}
}

func (s *loggingService) Register(user manager.User) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "register",
			"email", user.Email,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	return s.Service.Register(user)
}

func (s *loggingService) Login(user manager.User) (token string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "login",
			"email", user.Email,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	return s.Service.Login(user)
}
