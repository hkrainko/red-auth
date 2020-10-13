package domain

import "context"

type AuthType string


type AuthRepository interface {
	Auth(ctx context.Context, requester Requester) (Resident, error)
	GetAuthUrl(ctx context.Context, authType AuthType) (string, error)
}

type AuthUseCase interface {
	Auth(ctx context.Context, requester Requester) (Resident, error)
	GetAuthUrl(ctx context.Context, authType AuthType) (string, error)
}