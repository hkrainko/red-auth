package main

import (
	"context"
	"github.com/gin-gonic/gin"
	//"google.golang.org/grpc"
	//"log"
	//"net"
	pb "red-auth/proto"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *server) GetAuthUrl(ctx context.Context, request *pb.GetAuthUrlRequest) (*pb.GetAuthUrlResponse, error) {
	panic("implement me")
}

func main() {
	//lis, err := net.Listen("tcp", ":50051")
	//r := gin.Default()
	//
	//s := grpc.NewServer()
	//
	//pb.RegisterAuthServiceServer(s, InitAuthController())
	//
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	//err := r.Run(":9002")
	//print(err)

	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		ctr := InitAuthController()
		authGroup.POST("/getAuthUrl", ctr.GetAuthUrl)

	}

	err := r.Run(":9002")
	print(err)
}