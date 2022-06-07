package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"wallet_service/internal/endpoint"
	"wallet_service/internal/pkg"
	"wallet_service/internal/repository"
	"wallet_service/internal/service"
	"wallet_service/internal/transport"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
)

func main() {
	var httpAddr = flag.String("http", ":8085", "http listen address")

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "wallet",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	flag.Parse()
	ctx := context.Background()

	var srv service.Service
	{
		repository := repository.NewRepo(logger)
		srv = pkg.NewService(repository, logger)
	}

	errs := make(chan error)

	endpoints := endpoint.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := transport.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
