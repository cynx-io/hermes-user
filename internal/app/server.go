package app

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"hermes/internal/dependencies"
	"hermes/internal/grpc"
	"hermes/internal/pkg/logger"
	"strconv"
)

type Servers struct {
	grpcServer *grpc.Server
	config     *dependencies.Config
}

func (app *App) NewServers() (*Servers, error) {
	config := app.Dependencies.Config

	// Create gRPC server
	grpcServer := grpc.NewServer(app.Services.UserService)

	return &Servers{
		grpcServer: grpcServer,
		config:     config,
	}, nil
}

func (s *Servers) Start() error {
	var g errgroup.Group

	g.Go(func() error {
		logger.Infoln("Starting gRPC server")
		address := s.config.App.Address + ":" + strconv.Itoa(s.config.App.Port)
		if err := s.grpcServer.Start(address); err != nil {
			return fmt.Errorf("failed to start gRPC server: %w", err)
		}
		return nil
	})

	return g.Wait()
}

func (s *Servers) Stop() error {
	return errors.New("stop not implemented")
}
