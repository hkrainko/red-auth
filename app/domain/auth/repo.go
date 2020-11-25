package auth

import "context"

type Repo interface {
	HandleAuthCallBack(ctx context.Context, authCallBack CallBack) (error)
	GetAuthUrl(ctx context.Context, authType Type) (string, error)
}