package app

import (
	"hermes/internal/pkg/logger"
	"log"
)

type App struct {
	Dependencies *Dependencies
	Repos        *Repos
	Services     *Services
}

func NewApp(configPath string) (*App, error) {

	log.Println("Initializing Dependencies")
	dependencies := NewDependencies(configPath)

	if dependencies.Config.Database.AutoMigrate {
		logger.Infoln("Running database migrations")
		err := dependencies.DatabaseClient.RunMigrations()
		if err != nil {
			logger.Fatalln("Failed to run migrations: ", err)
		}
	}

	logger.Infoln("Initializing Repositories")
	repos := NewRepos(dependencies)

	logger.Infoln("Initializing Services")
	services := NewServices(repos, dependencies)

	logger.Infoln("App initialized")
	return &App{
		Dependencies: dependencies,
		Repos:        repos,
		Services:     services,
	}, nil
}
