package app

import (
	"github.com/cynxees/hermes-user/internal/repository/database"
)

type Repos struct {
	TblUser *database.TblUser
}

func NewRepos(dependencies *Dependencies) *Repos {

	return &Repos{
		TblUser: database.NewTblUser(dependencies.DatabaseClient.Db),
	}
}
