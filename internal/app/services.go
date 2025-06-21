package app

import (
	"github.com/cynxees/hermes-user/internal/module/usermodule"
)

type Services struct {
	UserService *usermodule.UserService
}

func NewServices(repos *Repos, dependencies *Dependencies) *Services {

	return &Services{
		UserService: usermodule.NewUserService(repos.TblUser),
	}
}
