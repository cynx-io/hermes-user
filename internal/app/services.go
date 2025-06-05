package app

import (
	"hermes/internal/module/usermodule"
)

type Services struct {
	UserService *usermodule.UserService
}

func NewServices(repos *Repos, dependencies *Dependencies) *Services {

	return &Services{
		UserService: usermodule.NewUserService(repos.TblUser),
	}
}
