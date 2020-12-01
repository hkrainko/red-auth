package auth

import (
	"context"
)

type Repo interface {
	GetAuthorizedUserInfo(ctx context.Context, authCallBack CallBack) (*UserInfo, error)
	GetAuthUrl(ctx context.Context, authType Type) (string, error)
}