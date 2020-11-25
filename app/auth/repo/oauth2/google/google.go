package google

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"red-auth/app/domain/auth"
)

var (
	OauthConfig *oauth2.Config
)

func init() {
	OauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:9002/auth/callback?auth_type=Google",
		ClientID:     "601833756814-1n1uo2jp77sp888mjsrsl1fmru69kvhb.apps.googleusercontent.com",
		ClientSecret: "yyZUseQoTSYS8tT0MP-M9MrL",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func GetUserInfo(state string, code string) (*auth.AuthorizedUserInfo, error) {
	if state != "pseudo-random" {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	fmt.Println(string(contents))
	googleUserInfo := googleAuthorizedUserInfo{}
	err = json.Unmarshal(contents, &googleUserInfo)
	if err != nil {
		return nil, err
	}

	return googleUserInfo.toAuthorizedUserInfo(), nil
}

type googleAuthorizedUserInfo struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	PhotoUrl string `json:"picture"`
}

func (g *googleAuthorizedUserInfo) toAuthorizedUserInfo() *auth.AuthorizedUserInfo {
	return &auth.AuthorizedUserInfo{
		ID:       g.ID,
		Email:    g.Email,
		PhotoUrl: g.PhotoUrl,
	}
}
