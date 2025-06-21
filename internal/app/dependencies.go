package app

import (
	"context"
	"github.com/cynxees/cynx-core/src/logger"
	"github.com/cynxees/hermes-user/internal/dependencies"
)

type Dependencies struct {
	DatabaseClient *dependencies.DatabaseClient
}

func NewDependencies(ctx context.Context) *Dependencies {

	logger.Info(ctx, "Connecting to Database")
	databaseClient, err := dependencies.NewDatabaseClient()
	if err != nil {
		logger.Fatal(ctx, "Failed to connect to database: ", err)
	}

	logger.Info(ctx, "Dependencies initialized")
	return &Dependencies{
		DatabaseClient: databaseClient,
	}
}
