package app

import (
	"hermes/internal/dependencies"
	"hermes/internal/pkg"
	"hermes/internal/pkg/logger"
	"log"
)

type Dependencies struct {
	Config *dependencies.Config

	DatabaseClient *dependencies.DatabaseClient
}

func NewDependencies(configPath string) *Dependencies {

	log.Println("Loading Config")
	config, err := dependencies.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	log.Println("Initializing Logger")
	logger.InitLogger()

	logger.Info("Initializing Validator")
	pkg.InitValidator()

	logger.Info("Connecting to Database")
	databaseClient, err := dependencies.NewDatabaseClient(config)
	if err != nil {
		logger.Fatal("Failed to connect to database: ", err)
	}

	logger.Info("Dependencies initialized")
	return &Dependencies{
		Config:         config,
		DatabaseClient: databaseClient,
	}
}
