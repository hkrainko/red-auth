package auth

import (
	"errors"
)

type Type string

const (
	AuthTypeFacebook = "Facebook"
	AuthTypeTwitter  = "Twitter"
	AuthTypeGoogle   = "Google"
)

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
