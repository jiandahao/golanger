package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jiandahao/golanger/cmd/protoc-gen-golangergin/example/account"
)

func main() {
	accountClient := account.NewAccountClient("http://127.0.0.1:8080", http.DefaultClient)

	profile, err := accountClient.GetProfile(context.Background(), &account.GetProfileRequest{
		UserId:     "1234",
		CreateTime: "12356",
		Token:      "sfsfsfsfdfsd",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", *profile)
}
