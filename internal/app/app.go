package app

import (
	"context"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/hermes-user/internal/dependencies/config"
)

type App struct {
	Dependencies *Dependencies
	Repos        *Repos
	UseCases     *UseCases
}

func NewApp(ctx context.Context) (*App, error) {

	logger.Info(ctx, "Initializing Dependencies")
	dependencies := NewDependencies(ctx)

	if config.Config.Database.AutoMigrate {
		logger.Info(ctx, "Running database migrations")
		err := dependencies.DatabaseClient.RunMigrations(ctx)
		if err != nil {
			logger.Fatal(ctx, "Failed to run migrations: ", err)
		}
	}

	logger.Info(ctx, "Initializing Repositories")
	repos := NewRepos(dependencies)

	logger.Info(ctx, "Initializing Use Cases")
	useCases := NewUseCases(repos, dependencies)

	logger.Info(ctx, "App initialized")
	return &App{
		Dependencies: dependencies,
		Repos:        repos,
		UseCases:     useCases,
	}, nil
}
