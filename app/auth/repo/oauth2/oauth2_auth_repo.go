package oauth2

import (
	"context"
	"red-auth/domain/auth"
	"red-auth/domain/auth/model"
)


type oauth2AuthRepo struct {
	OAuth2Configs map[auth.Type]model.OAuth2
}

func NewOAuth2AuthRepository(oAuth2Configs map[auth.Type]model.OAuth2) auth.Repo {
	return &oauth2AuthRepo{
		oAuth2Configs,
	}
}

func (o oauth2AuthRepo) GetAuthUrl(ctx context.Context, authType auth.Type) (string, error) {
	if config, ok := o.OAuth2Configs[authType]; ok {
		return config.GetAuthUrl(), nil
	}
	return "", auth.NewAuthTypeNotFoundError()
}

func (o oauth2AuthRepo) GetAuthorizedUserInfo(ctx context.Context, authCallBack auth.CallBack) (*auth.UserInfo, error) {
	if config, ok := o.OAuth2Configs[authCallBack.AuthType]; ok {
		return config.GetUserInfo(authCallBack.State, authCallBack.Code)
	}
	return nil, auth.NewAuthTypeNotFoundError()
}

