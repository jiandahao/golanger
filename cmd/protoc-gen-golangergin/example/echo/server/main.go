package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jiandahao/golanger/cmd/protoc-gen-golangergin/example/echo"
)

type echoServer struct {
	echo.UnimplementedEchoServer
}

func (s *echoServer) GetEcho(ctx context.Context, in *echo.GetEchoReq) (*echo.GetEchoResp, error) {
	return &echo.GetEchoResp{
		ParamInUriOrQuery:    in.ParamInUriOrQuery,
		ParamInHeaderOrQuery: in.ParamInHeaderOrQuery,
	}, nil
}

func (s *echoServer) PostEcho(ctx context.Context, in *echo.PostEchoReq) (*echo.PostEchoResp, error) {
	return &echo.PostEchoResp{
		ParamInUriOrQuery: in.ParamInUriOrQuery,
		ParamInHeader:     in.ParamInHeader,
		ParamInBody:       in.ParamInBody,
	}, nil
}

func (s *echoServer) PostFormEcho(ctx context.Context, in *echo.PostFormEchoReq) (*echo.PostFormEchoResp, error) {
	var filenameA, filenameB string
	for _, f := range in.FilesA {
		filenameA += f.Filename + ","
	}

	if in.FileB != nil {
		filenameB = in.FileB.Filename
	}
	return &echo.PostFormEchoResp{
		ParamInFormA: in.ParamInFormA,
		ParamInFormB: in.ParamInFormB,
		FilenameA:    filenameA,
		FilenameB:    filenameB,
	}, nil
}

func main() {
	router := gin.New()

	echo.RegisterEchoServer(router, &echoServer{})

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
		return
	}
}
