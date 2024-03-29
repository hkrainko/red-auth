// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"red-auth/app/auth/delivery/gRPC"
	"red-auth/app/auth/repo/oauth2"
	"red-auth/app/auth/usecase"
	"red-auth/domain/auth"
	"red-auth/domain/auth/model"
)

// Injectors from wire.go:

func InitAuthController(oAuth2 map[auth.Type]model.OAuth2) gRPC.AuthController {
	repo := oauth2.NewOAuth2AuthRepository(oAuth2)
	useCase := usecase.NewAuthUseCase(repo)
	authController := gRPC.NewAuthController(useCase)
	return authController
}
