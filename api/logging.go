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

func (ls *loggingService) Register(user manager.User) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			"method", "register",
			"email", user.Email,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	return ls.Service.Register(user)
}

func (ls *loggingService) Login(user manager.User) (token string, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			"method", "login",
			"email", user.Email,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	return ls.Service.Login(user)
}

func (ls *loggingService) AddClient(key string, client manager.Client) (id string, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			"method", "add_client",
			"owner", client.Owner,
			"id", id,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	return ls.Service.AddClient(key, client)
}

func (ls *loggingService) ViewClient(key string, id string) (client manager.Client, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			"method", "view_client",
			"owner", client.Owner,
			"id", id,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	return ls.Service.ViewClient(key, id)
}

func (ls *loggingService) RemoveClient(key string, id string) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			"method", "remove_client",
			"id", id,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	return ls.Service.RemoveClient(key, id)
}
