package http

import (
	"github.com/gin-gonic/gin"
	"red-auth/app/auth/delivery/get-auth-url"
	"red-auth/app/domain"
)

type AuthController struct{
	authUsecase domain.AuthUseCase
}

func NewAuthController(useCase domain.AuthUseCase) AuthController {
	return AuthController{
		authUsecase: useCase,
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
	c.JSON(200, get_auth_url.NewResponse(url))
}