package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/mainflux/manager"
	"github.com/mainflux/manager/api"
	"github.com/mainflux/manager/mocks"
)

func main() {
	users := mocks.NewUserRepositoryMock()
	idp := mocks.NewIdentityProviderMock()

	var logger log.Logger
	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var svc manager.Service
	svc = manager.NewService(users, idp)
	svc = api.NewLoggingService(logger, svc)

	handler := api.MakeHandler(svc)

	http.ListenAndServe(":8180", handler)
}
