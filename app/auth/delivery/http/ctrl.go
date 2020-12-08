package http

import (
	"github.com/gin-gonic/gin"
	call_back "red-auth/app/auth/delivery/call-back"
	"red-auth/app/auth/delivery/get-auth-url"
	"red-auth/domain/auth"
)

type AuthController struct{
	authUseCase auth.UseCase
}

func NewAuthController(useCase auth.UseCase) AuthController {
	return AuthController{
		authUseCase: useCase,
	}

}

func (a AuthController) GetAuthUrl(c *gin.Context) {
	authType := c.PostForm("auth_type")
	if authType == "" {
		return
	}
	url, err := a.authUseCase.GetAuthUrl(c, auth.Type(authType))
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
	userInfo, err := a.authUseCase.HandleAuthCallBack(c, auth.CallBack{
		AuthType: auth.Type(authType),
		State:    state,
		Code:     code,
	})
	if err != nil {
		return
	}

	c.JSON(200, call_back.NewResponse(userInfo))
}