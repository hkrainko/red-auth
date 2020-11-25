package usecase

import (
	"context"
	"red-auth/app/domain/auth"
)

type authUsecase struct {
	authRepo auth.Repo
}

func NewAuthUseCase(repo auth.Repo) auth.UseCase {
	return &authUsecase {
		authRepo: repo,
	}
}

func (a authUsecase) HandleAuthCallBack(ctx context.Context, authCallBack auth.CallBack) (error) {
	return a.authRepo.HandleAuthCallBack(ctx, authCallBack)
}

func (a authUsecase) GetAuthUrl(ctx context.Context, authType auth.Type) (string, error) {
	return a.authRepo.GetAuthUrl(ctx, authType)
}
