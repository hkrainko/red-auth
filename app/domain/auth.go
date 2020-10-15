package domain

import (
	"context"
	"errors"
)

type AuthType string


type AuthRepository interface {
	Auth(ctx context.Context, requester Requester) (Resident, error)
	GetAuthUrl(ctx context.Context, authType AuthType) (string, error)
}

type AuthUseCase interface {
	Auth(ctx context.Context, requester Requester) (Resident, error)
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
			msg: "AuthTypeNotFoundError",
			err: errors.New(""),
		},
	}
}




