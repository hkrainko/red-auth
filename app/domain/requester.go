package domain

import "context"

type Requester struct {
	AuthType string
	AuthURL string
}

type RequesterRepository interface {
	GetAuthURL(ctx context.Context, authType AuthType) string
}