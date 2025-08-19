package app

import (
	"github.com/cynx-io/hermes-user/internal/repository/database"
)

type Repos struct {
	UserRepo *database.UserRepo
}

func NewRepos(dependencies *Dependencies) *Repos {
	return &Repos{
		UserRepo: database.NewUserRepo(dependencies.DatabaseClient.Db),
	}
}
