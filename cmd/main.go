package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/mainflux/manager"
	"github.com/mainflux/manager/api"
	"github.com/mainflux/manager/mocks"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var opts struct {
	Port int
}

func main() {
	flag.IntVar(&opts.Port, "port", 8180, "HTTP server port")
	flag.Parse()

	users := mocks.NewUserRepositoryMock()
	idp := mocks.NewIdentityProviderMock()

	var logger log.Logger
	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var fields = []string{"method"}

	var svc manager.Service
	svc = manager.NewService(users, idp)
	svc = api.NewLoggingService(logger, svc)
	svc = api.NewMetricService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "manager",
			Subsystem: "api",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fields),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "manager",
			Subsystem: "api",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fields),
		svc,
	)

	errs := make(chan error, 2)

	go func() {
		p := fmt.Sprintf(":%d", opts.Port)
		errs <- http.ListenAndServe(p, api.MakeHandler(svc))
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}
