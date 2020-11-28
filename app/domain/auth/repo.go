package auth

import "context"

type Repo interface {
	GetAuthorizedUserInfo(ctx context.Context, authCallBack CallBack) (*AuthorizedUserInfo, error)
	GetAuthUrl(ctx context.Context, authType Type) (string, error)
}