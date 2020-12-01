package token

import "context"

type Repo interface {
	GenerateAuthToken(ctx context.Context) (string, error)
}