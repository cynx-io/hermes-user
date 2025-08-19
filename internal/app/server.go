package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/hermes-user/internal/dependencies/config"
	"github.com/cynx-io/hermes-user/internal/grpc"
	"golang.org/x/sync/errgroup"
	"strconv"
)

type Servers struct {
	grpcServer *grpc.Server
}

func (app *App) NewServers() (*Servers, error) {
	// Create gRPC server
	grpcServer := grpc.NewServer(app.UseCases.UserUseCase)

	return &Servers{
		grpcServer: grpcServer,
	}, nil
}

func (s *Servers) Start(ctx context.Context) error {
	var g errgroup.Group

	g.Go(func() error {
		logger.Info(ctx, "Starting gRPC server")
		address := config.Config.App.Address + ":" + strconv.Itoa(config.Config.App.Port)
		if err := s.grpcServer.Start(ctx, address); err != nil {
			return fmt.Errorf("failed to start gRPC server: %w", err)
		}
		return nil
	})

	return g.Wait()
}

func (s *Servers) Stop() error {
	return errors.New("stop not implemented")
}
