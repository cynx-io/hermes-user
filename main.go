package main

import (
	"context"
	"github.com/cynxees/cynx-core/src/logger"
	"github.com/cynxees/hermes-user/internal/app"
	"github.com/cynxees/hermes-user/internal/dependencies/config"
	"github.com/sirupsen/logrus"
)

func main() {

	ctx := context.Background()
	defer ctx.Done()

	config.Init()
	logLevel, err := logrus.ParseLevel(config.Config.Elastic.Level)
	if err != nil {
		logLevel = logrus.DebugLevel
	}
	logger.Init(logger.LoggerConfig{
		Level:            logLevel,
		ElasticsearchURL: []string{config.Config.Elastic.Url},
		ServiceName:      "hermes-user",
	})

	logger.Info(ctx, "Starting hermes")
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	logger.Info(ctx, "Initializing App")
	application, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}

	logger.Info(ctx, "Creating servers")
	servers, err := application.NewServers()
	if err != nil {
		panic(err)
	}

	logger.Info(ctx, "Starting servers")
	if err := servers.Start(ctx); err != nil {
		panic(err)
	}
}
