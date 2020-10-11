package oauth2

import "red-auth/app/domain"


type oauth2AuthRepo struct {

}

func NewOAuth2AuthRepository() domain.AuthRepository {
	return &oauth2AuthRepo{
		
	}
}

func (o oauth2AuthRepo) GetAuthUrl(authType domain.AuthType) string {
	panic("implement me")
}

func (o oauth2AuthRepo) Auth(requester domain.Requester) (domain.Resident, error) {
	panic("implement me")

	//TODO: Redirect


}