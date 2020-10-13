package usecase

import (
	"context"
	"red-auth/app/domain"
)

type authUsecase struct {
	authRepo domain.AuthRepository
}

func NewArtistUsecase(repo domain.AuthRepository) domain.AuthUseCase {
	return &authUsecase{
		authRepo: repo,
	}
}

func (a authUsecase) Auth(ctx context.Context, requester domain.Requester) (domain.Resident, error) {
	panic("implement me")
}

func (a authUsecase) GetAuthUrl(ctx context.Context, authType domain.AuthType) (string, error) {
	panic("implement me")
}

