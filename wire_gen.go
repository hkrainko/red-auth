// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"red-auth/app/auth/delivery/gRPC"
	"red-auth/app/auth/repository/oauth2"
	"red-auth/app/auth/usecase"
)

// Injectors from wire.go:

func InitAuthController() gRPC.AuthController {
	authRepository := oauth2.NewOAuth2AuthRepository()
	authUseCase := usecase.NewAuthUsecase(authRepository)
	authController := gRPC.NewAuthController(authUseCase)
	return authController
}
