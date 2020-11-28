//+build wireinject

package main

import (
	"github.com/google/wire"
	auth_deliv "red-auth/app/auth/delivery/gRPC"
	auth_repo "red-auth/app/auth/repo/oauth2"
	auth_ucase "red-auth/app/auth/usecase"
)

func InitAuthController() auth_deliv.AuthController {
	wire.Build(auth_deliv.NewAuthController, auth_ucase.NewAuthUseCase, auth_repo.NewOAuth2AuthRepository)
	return auth_deliv.AuthController{}
}