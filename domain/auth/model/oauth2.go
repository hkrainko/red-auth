package model

import "red-auth/domain/auth"

type OAuth2 interface {
	GetUserInfo(state string, code string) (*auth.UserInfo, error)
	GetAuthUrl() string
}