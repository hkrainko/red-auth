package facebook

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"io/ioutil"
	"net/http"
	"red-auth/app/domain/auth"
)

var (
	OauthConfig *oauth2.Config
)

func init() {
	OauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:9002/auth/callback?auth_type=Facebook",
		ClientID:     "375183940527894",
		ClientSecret: "d90ae5b9183d30d3cc206c2286b15e90",
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
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
	response, err := http.Get("https://graph.facebook.com/v9.0/me/picture?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	fmt.Println(string(contents))
	facebookUserInfo := facebookAuthorizedUserInfo{}
	err = json.Unmarshal(contents, &facebookUserInfo)
	return facebookUserInfo.toAuthorizedUserInfo(), nil
}

type facebookAuthorizedUserInfo struct {
	ID       string `json:"id"`
	Name    string `json:"name"`
	PhotoUrl string `json:"picture"`
}

func (g *facebookAuthorizedUserInfo) toAuthorizedUserInfo() *auth.AuthorizedUserInfo {
	return &auth.AuthorizedUserInfo{
		ID:       g.ID,
		PhotoUrl: g.PhotoUrl,
	}
}