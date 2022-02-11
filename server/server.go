package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

type Server struct {
	logger     *log.Logger
	httpServer *fasthttp.Server
	errorChan  chan error
}

func NewServer(logger *log.Logger, router *fasthttprouter.Router) *Server {
	return &Server{
		logger:     logger,
		httpServer: &fasthttp.Server{Handler: router.Handler},
		errorChan:  make(chan error, 1),
	}
}

func (s *Server) Start(ctx context.Context, listenAddr string) error {
	go func(errorChan chan error) {
		errorChan <- s.httpServer.ListenAndServe(listenAddr)
	}(s.errorChan)

	defer s.backgroundErrorHandling()

	select {
	case <-ctx.Done():
		return errors.WithStack(ctx.Err())
	case err := <-s.errorChan:
		return errors.WithStack(err)
	case <-time.NewTimer(time.Second).C:
		return nil
	}
}

func (s *Server) backgroundErrorHandling() {
	go func(logger *log.Logger, errorChan chan error) {
		for err := range errorChan {
			if err == nil {
				continue
			}

			logger.Fatal(context.Background(), fmt.Sprintf("server failed: %s", errors.WithStack(err)))
		}
	}(s.logger, s.errorChan)
}

func (s *Server) Shutdown() error {
	return s.httpServer.Shutdown()
}

func RunServer(lc fx.Lifecycle, logger *log.Logger, config *Config, srv *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			addr := net.JoinHostPort(config.APIBindingAddr, config.APIBindingPort)
			logger.Println(fmt.Sprintf("Running server on %s", addr))

			return srv.Start(ctx, addr)
		},
		OnStop: func(_ context.Context) error {
			logger.Println("Server shutting down...")

			return srv.Shutdown()
		},
	})
}
