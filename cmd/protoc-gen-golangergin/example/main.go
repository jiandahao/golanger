package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jiandahao/golanger/cmd/protoc-gen-golangergin/example/account"
)

type accountServer struct {
	account.UnimplementedAccountServer
}

func (s *accountServer) CreateAccount(ctx context.Context, in *account.AccountRegister) (*account.RegisterStatus, error) {
	fmt.Println(in.Username, in.Email)
	return &account.RegisterStatus{Status: "success"}, nil
}
func (s *accountServer) GetProfile(ctx context.Context, in *account.GetProfileRequest) (*account.Profile, error) {
	fmt.Println("====", in.UserId, in.CreateTime)

	return &account.Profile{UserId: in.UserId, Username: "jian"}, nil
}

func main() {
	router := gin.New()

	account.RegisterAccountServer(router, &accountServer{})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
