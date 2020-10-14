package gRPC

import (
	"context"
	"red-auth/app/domain"
	pb "red-auth/proto"
)

type AuthController struct{
	authUsecase domain.AuthUseCase
	pb.UnimplementedAuthServiceServer
}

func NewAuthController(usecase domain.AuthUseCase) AuthController {
	return AuthController{
		authUsecase: usecase,
	}
}

func (a AuthController) GetAuthUrl(ctx context.Context, request *pb.GetAuthUrlRequest) (*pb.GetAuthUrlResponse, error) {
	url, err := a.authUsecase.GetAuthUrl(ctx, domain.AuthType(request.AuthType))
	if err != nil {
		return nil, err
	}
	return &pb.GetAuthUrlResponse{
		AuthUrl: url,
	}, nil
}

