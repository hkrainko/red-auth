package google

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"red-auth/domain/auth"
)

type OAuth struct {
	OauthConfig *oauth2.Config
}

func InitGoogleOAuth(redirectURL string, clientID string, clientSecret string) OAuth {
	return OAuth {
		OauthConfig: &oauth2.Config{
			RedirectURL:  redirectURL,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
				"https://www.googleapis.com/auth/user.birthday.read",
				"https://www.googleapis.com/auth/user.gender.read",
			},
			Endpoint: google.Endpoint,
		},
	}
}

func (o OAuth)GetUserInfo(state string, code string) (*auth.UserInfo, error) {
	if state != "pseudo-random" {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := o.OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	fmt.Println(token)
	response, err := http.Get("https://people.googleapis.com/v1/people/me?personFields=names,birthdays,photos,genders,emailAddresses&access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	fmt.Println(string(contents))
	googleUserInfo := AuthorizedUserInfo{}
	err = json.Unmarshal(contents, &googleUserInfo)
	if err != nil {
		return nil, err
	}

	return googleUserInfo.toUserInfo()
}

func (o OAuth) GetAuthUrl() string {
	return o.OauthConfig.AuthCodeURL("pseudo-random")
}