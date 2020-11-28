package usecase

import (
	"context"
	"red-auth/app/domain/auth"
)

type authUseCase struct {
	authRepo auth.Repo
}

func NewAuthUseCase(repo auth.Repo) auth.UseCase {
	return &authUseCase{
		authRepo: repo,
	}
}

func (a authUseCase) HandleAuthCallBack(ctx context.Context, authCallBack auth.CallBack) (*auth.AuthorizedUserInfo, error) {
	userInfo, err := a.authRepo.GetAuthorizedUserInfo(ctx, authCallBack)
	if err != nil {
		return nil, err
	}
	//TODO: Event for new user
	return userInfo, nil
}

func (a authUseCase) GetAuthUrl(ctx context.Context, authType auth.Type) (string, error) {
	return a.authRepo.GetAuthUrl(ctx, authType)
}
