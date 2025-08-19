package app

import (
	"github.com/cynx-io/hermes-user/internal/usecase/userusecase"
)

type UseCases struct {
	UserUseCase *userusecase.UseCase
}

func NewUseCases(repos *Repos, dependencies *Dependencies) *UseCases {
	return &UseCases{
		UserUseCase: userusecase.NewUseCase(repos.UserRepo),
	}
}
