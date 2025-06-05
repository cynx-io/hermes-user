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
	logger.InitLogger(&config.Log)

	logger.Infoln("Initializing Validator")
	pkg.InitValidator()

	logger.Infoln("Connecting to Database")
	databaseClient, err := dependencies.NewDatabaseClient(config)
	if err != nil {
		logger.Fatalln("Failed to connect to database: ", err)
	}

	logger.Infoln("Dependencies initialized")
	return &Dependencies{
		Config:         config,
		DatabaseClient: databaseClient,
	}
}
