#!/bin/bash
set -x

export GO111MODULE=on
go version && go env

mkdir bin
go build -o bin/pprof_proxy cmd/pprof_proxy/main.go
go build -o bin/gen_swagger cmd/gen_swagger/main.go
go build -o bin/golanger cmd/golanger/main.go