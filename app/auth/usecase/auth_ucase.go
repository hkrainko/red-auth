package usecase

import (
	"context"
	"red-auth/app/domain"
)

type authUsecase struct {
	authRepo domain.AuthRepository
}

func NewAuthUseCase(repo domain.AuthRepository) domain.AuthUseCase {
	return &authUsecase{
		authRepo: repo,
	}
}

func (a authUsecase) HandleAuthCallBack(ctx context.Context, authCallBack domain.AuthCallBack) (error) {
	return a.authRepo.HandleAuthCallBack(ctx, authCallBack)
}

func (a authUsecase) GetAuthUrl(ctx context.Context, authType domain.AuthType) (string, error) {
	return a.authRepo.GetAuthUrl(ctx, authType)
}
