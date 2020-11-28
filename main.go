package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"

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
	lis, err := net.Listen("tcp", ":50051")
	r := gin.Default()

	s := grpc.NewServer()

	pb.RegisterAuthServiceServer(s, InitAuthController())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = r.Run(":9002")
	fmt.Println(err)

	//r := gin.Default()
	//
	//authGroup := r.Group("/auth")
	//{
	//	ctr := InitAuthController()
	//	authGroup.POST("/getAuthUrl", ctr.GetAuthUrl)
	//	authGroup.GET("/callback", ctr.CallBack)
	//}
	//
	//err := r.Run(":9002")
	//print(err)
}