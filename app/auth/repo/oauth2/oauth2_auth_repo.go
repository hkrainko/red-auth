package oauth2

import (
	"context"
	"red-auth/app/auth/repo/oauth2/facebook"
	"red-auth/app/auth/repo/oauth2/google"
	"red-auth/domain/auth"
)


type oauth2AuthRepo struct {

}

func NewOAuth2AuthRepository() auth.Repo {
	return &oauth2AuthRepo{
		
	}
}

func (o oauth2AuthRepo) GetAuthUrl(ctx context.Context, authType auth.Type) (string, error) {
	//return fmt.Sprintf("red-auth:%v", string(authType)), nil

	switch authType {
	case "Google":
		url := google.OauthConfig.AuthCodeURL("pseudo-random")
		return url, nil
	case "Facebook":
		url := facebook.OauthConfig.AuthCodeURL("pseudo-random")
		return url, nil
	default:
		break
	}
	return "", auth.NewAuthTypeNotFoundError()
}

func (o oauth2AuthRepo) GetAuthorizedUserInfo(ctx context.Context, authCallBack auth.CallBack) (*auth.UserInfo, error) {
	switch authCallBack.AuthType {
	case "Google":
		return google.GetUserInfo(authCallBack.State, authCallBack.Code)
	case "Facebook":
		return facebook.GetUserInfo(authCallBack.State, authCallBack.Code)
	default:
		break
	}
	return nil, auth.NewAuthTypeNotFoundError()
}

