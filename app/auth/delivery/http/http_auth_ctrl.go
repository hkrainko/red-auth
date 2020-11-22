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
	authType := c.PostForm("auth_type")
	if authType == "" {
		return
	}
	url, err := a.authUsecase.GetAuthUrl(c, domain.AuthType(authType))
	if err != nil {
		return
	}
	c.JSON(200, get_auth_url.NewResponse(url))
}

func (a AuthController) CallBack(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")
	authType := c.Query("auth_type")
	if state == "" || code == "0" {
		return
	}
	err := a.authUsecase.HandleAuthCallBack(c, domain.AuthCallBack{
		AuthType: domain.AuthType(authType),
		State: state,
		Code:  code,
	})
	if err != nil {
		return
	}
	//c.JSON(200, get_auth_url.NewResponse(url))
}