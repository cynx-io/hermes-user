package app

import (
	"context"
	"github.com/cynxees/cynx-core/src/logger"
	"github.com/cynxees/hermes-user/internal/dependencies/config"
)

type App struct {
	Dependencies *Dependencies
	Repos        *Repos
	Services     *Services
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

	logger.Info(ctx, "Initializing Services")
	services := NewServices(repos, dependencies)

	logger.Info(ctx, "App initialized")
	return &App{
		Dependencies: dependencies,
		Repos:        repos,
		Services:     services,
	}, nil
}
