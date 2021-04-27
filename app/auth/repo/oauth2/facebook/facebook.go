package facebook

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"io/ioutil"
	"net/http"
	"red-auth/domain/auth"
)

type OAuth struct {
	OauthConfig *oauth2.Config
}

func InitFacebookOAuth(redirectURL string, clientID string, clientSecret string) OAuth {
	return OAuth{
		OauthConfig: &oauth2.Config{
			RedirectURL:  redirectURL,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{"public_profile", "user_birthday", "email", "user_gender"},
			Endpoint:     facebook.Endpoint,
		},
	}
}

func (o OAuth) GetUserInfo(state string, code string) (*auth.UserInfo, error) {
	if state != "pseudo-random" {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := o.OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://graph.facebook.com/v9.0/me?fields=id,name,birthday,email,gender,picture&access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	fmt.Println(string(contents))
	facebookUserInfo := AuthorizedUserInfo{}
	err = json.Unmarshal(contents, &facebookUserInfo)
	return facebookUserInfo.toUserInfo(), nil
}

func (o OAuth) GetAuthUrl() string {
	return o.OauthConfig.AuthCodeURL("pseudo-random")
}