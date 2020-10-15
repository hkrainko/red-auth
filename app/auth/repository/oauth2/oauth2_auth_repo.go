package oauth2

import (
	"context"
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
		url := googleOauthConfig.AuthCodeURL("pseudo-random")
		return url, nil
	default:
		break
	}
	return "", domain.NewAuthTypeNotFoundError()
}

func (o oauth2AuthRepo) Auth(ctx context.Context, requester domain.Requester) (domain.Resident, error) {
	panic("implement me")

	//TODO: Redirect


}