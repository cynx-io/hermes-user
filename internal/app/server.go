package app

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"hermes/internal/grpc"
	"hermes/internal/pkg/logger"
	"strconv"
)

type Servers struct {
	grpcServer *grpc.Server
}

func (app *App) NewServers() (*Servers, error) {
	config := app.Dependencies.Config

	// Create gRPC server
	grpcServer := grpc.NewServer(app.Services.UserService)

	address := config.App.Address + ":" + strconv.Itoa(config.App.Port)
	logger.Infof("Starting gRPC server on %s", address)

	return &Servers{
		grpcServer: grpcServer,
	}, nil
}

func (s *Servers) Start() error {
	var g errgroup.Group

	g.Go(func() error {
		logger.Infoln("Starting gRPC server")
		if err := s.grpcServer.Start(":5000"); err != nil {
			return fmt.Errorf("failed to start gRPC server: %w", err)
		}
		return nil
	})

	return g.Wait()
}

func (s *Servers) Stop() error {
	return errors.New("stop not implemented")
}
