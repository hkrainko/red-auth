package gRPC

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"red-auth/app/domain"
)

type AuthController struct{
	grpcConn *grpc.ClientConn
	authUsecase domain.AuthUseCase
}

func NewAuthController(grpcConn *grpc.ClientConn, usecase domain.AuthUseCase) AuthController {
	return AuthController{
		grpcConn: grpcConn,
		authUsecase: usecase,
	}
}

func (a AuthController) GetAuthUrl(c *gin.Context) {
	authType := c.PostForm("authType")
	if authType == "" {
		return
	}
	url, err := a.authUsecase.GetAuthUrl(c, domain.AuthType(authType))
	if err != nil {
		return
	}



}

