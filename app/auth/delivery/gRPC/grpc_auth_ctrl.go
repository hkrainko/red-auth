package gRPC

import (
	"context"
	"red-auth/app/domain/auth"
	pb "red-auth/proto"
)

type AuthController struct{
	authUseCase auth.UseCase
	pb.UnimplementedAuthServiceServer
}

func NewAuthController(useCase auth.UseCase) AuthController {
	return AuthController{
		authUseCase: useCase,
	}
}

func (a AuthController) GetAuthUrl(ctx context.Context, request *pb.GetAuthUrlRequest) (*pb.GetAuthUrlResponse, error) {
	url, err := a.authUseCase.GetAuthUrl(ctx, auth.Type(request.AuthType))
	if err != nil {
		return nil, err
	}
	return &pb.GetAuthUrlResponse{
		AuthUrl: url,
	}, nil
}

