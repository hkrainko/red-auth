package auth

import "context"

type UseCase interface {
	HandleAuthCallBack(ctx context.Context, authCallBack CallBack) (error)
	GetAuthUrl(ctx context.Context, authType Type) (string, error)
}