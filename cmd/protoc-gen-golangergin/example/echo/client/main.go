package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jiandahao/golanger/cmd/protoc-gen-golangergin/example/echo"
)

func main() {
	s := echo.NewEchoClient("http://127.0.0.1:8080", http.DefaultClient)

	{
		resp, err := s.GetEcho(context.Background(), &echo.GetEchoReq{
			ParamInUriOrQuery:    "1234",
			ParamInHeaderOrQuery: "2345",
		})

		if err != nil {
			panic(err)
		}

		fmt.Println(toJSONStr(resp))
	}

	{
		resp, err := s.PostEcho(context.Background(), &echo.PostEchoReq{
			ParamInUriOrQuery: "123",
			ParamInHeader:     "456",
			ParamInBody:       "789",
		})
		if err != nil {
			panic(err)
		}

		fmt.Println(toJSONStr(resp))
	}

	{
		resp, err := s.PostFormEcho(context.Background(), &echo.PostFormEchoReq{
			ParamInFormA: "1111",
			ParamInFormB: "2222",
			MultipartFiles: map[string]map[string]io.Reader{
				"files_a": {
					"filename1": bytes.NewBufferString("file bbbb"),
					"filename2": bytes.NewBufferString("file cccc"),
				},
				"file_b": {
					"filename2": bytes.NewBufferString("file bbbb"),
				},
			},
		})

		if err != nil {
			panic(err)
		}

		fmt.Println(toJSONStr(resp))
	}
}

func toJSONStr(v interface{}) string {
	data, _ := json.Marshal(v)

	return string(data)
}
