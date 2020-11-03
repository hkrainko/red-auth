package oauth2

import (
	"context"
	"red-auth/app/auth/repository/oauth2/google"
	"red-auth/app/domain"
)


type oauth2AuthRepo struct {

}

func NewOAuth2AuthRepository() domain.AuthRepository {
	return &oauth2AuthRepo{
		
	}
}

func (o oauth2AuthRepo) GetAuthUrl(ctx context.Context, authType domain.AuthType) (string, error) {
	//return fmt.Sprintf("red-auth:%v", string(authType)), nil

	switch authType {
	case "Google":
		url := google.GoogleOauthConfig.AuthCodeURL("pseudo-random")
		return url, nil
	default:
		break
	}
	return "", domain.NewAuthTypeNotFoundError()
}

func (o oauth2AuthRepo) HandleAuthCallBack(ctx context.Context, authCallBack domain.AuthCallBack) (error) {
	switch authCallBack.AuthType {
	case "Google":
		google.GetUserInfo(authCallBack.State, authCallBack.Code)
	default:
		break
	}
	return domain.NewAuthTypeNotFoundError()
}