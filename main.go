package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"red-auth/app/auth/repo/oauth2/facebook"
	"red-auth/app/auth/repo/oauth2/google"
	"red-auth/domain/auth"
	"red-auth/domain/auth/model"
	pb "red-auth/proto"
)

func init() {
	_ = os.MkdirAll("./tmp", 0777)
	f, err := os.OpenFile("./tmp/example.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(f)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	r := gin.Default()

	s := grpc.NewServer()

	facebookOAuthConfig := facebook.InitFacebookOAuth(
		"http://localhost:3000/auth?auth_type=Facebook",
		"375183940527894",
		"d90ae5b9183d30d3cc206c2286b15e90",
	)
	googleOAuthConfig := google.InitGoogleOAuth(
		"http://localhost:3000/auth?auth_type=Google",
		"601833756814-1n1uo2jp77sp888mjsrsl1fmru69kvhb.apps.googleusercontent.com",
		"yyZUseQoTSYS8tT0MP-M9MrL",
	)
	oAuthConfigs := map[auth.Type]model.OAuth2 {
		auth.AuthTypeFacebook: facebookOAuthConfig,
		auth.AuthTypeGoogle: googleOAuthConfig,
	}

	pb.RegisterAuthServiceServer(s, InitAuthController(oAuthConfigs))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = r.Run(":9002")
	fmt.Println(err)
}
