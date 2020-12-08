package gRPC

import (
	"context"
	"red-auth/domain/auth"
	pb "red-auth/proto"
)

type AuthController struct {
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

func (a AuthController) CallBack(ctx context.Context, request *pb.CallBackRequest) (*pb.CallBackResponse, error) {
	info, err := a.authUseCase.HandleAuthCallBack(ctx, auth.CallBack{
		AuthType: auth.Type(request.AuthType),
		State:    request.State,
		Code:     request.Code,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CallBackResponse{
		Id:       info.ID,
		AuthType: info.AuthType,
		Name:     info.Name,
		Email:    info.Email,
		Birthday: info.Birthday.Format("20060102"),
		Gender:   info.Gender,
		PhotoUrl: info.PhotoURL,
	}, nil
}
