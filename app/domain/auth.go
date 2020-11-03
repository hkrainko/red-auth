package domain

import (
	"context"
	"errors"
)

type AuthType string


type AuthRepository interface {
	HandleAuthCallBack(ctx context.Context, authCallBack AuthCallBack) (error)
	GetAuthUrl(ctx context.Context, authType AuthType) (string, error)
}

type AuthUseCase interface {
	HandleAuthCallBack(ctx context.Context, authCallBack AuthCallBack) (error)
	GetAuthUrl(ctx context.Context, authType AuthType) (string, error)
}

type AuthError struct {
	msg string
	err error
}

func (e *AuthError) Error() string {
	return e.err.Error() + e.msg
}

func (e *AuthError) Unwrap() string {
	return e.err.Error() + e.msg
}

type AuthTypeNotFoundError struct {
	AuthError
}

func NewAuthTypeNotFoundError() *AuthTypeNotFoundError {
	return &AuthTypeNotFoundError{
		AuthError{
			msg: "authTypeNotFoundError",
			err: errors.New("auth error"),
		},
	}
}




