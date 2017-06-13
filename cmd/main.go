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
	"github.com/mainflux/manager/cockroachdb"
	"github.com/mainflux/manager/mocks"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	port int    = 8180
	addr string = "postgresql://mainflux@0.0.0.0:26257/manager?sslmode=disable"
)

type flags struct {
	Port int
	Addr string
}

func main() {
	var cfg flags
	flag.IntVar(&cfg.Port, "port", port, "HTTP server port")
	flag.StringVar(&cfg.Addr, "db", addr, "database connection string")
	flag.Parse()

	var logger log.Logger
	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var fields = []string{"method"}

	db, err := cockroachdb.Connect(cfg.Addr)
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()

	users := cockroachdb.NewUserRepository(db)
	idp := mocks.NewIdentityProvider()

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
		p := fmt.Sprintf(":%d", cfg.Port)
		errs <- http.ListenAndServe(p, api.MakeHandler(svc))
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}
