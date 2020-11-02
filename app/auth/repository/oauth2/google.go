package oauth2

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     "601833756814-1n1uo2jp77sp888mjsrsl1fmru69kvhb.apps.googleusercontent.com",
		ClientSecret: "yyZUseQoTSYS8tT0MP-M9MrL",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}


