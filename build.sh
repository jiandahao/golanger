#!/bin/bash
set -x

export GO111MODULE=on
go version && go env

mkdir -p bin
go build -o bin/pprof_proxy cmd/pprof_proxy/main.go
go build -o bin/gen_swagger cmd/gen_swagger/main.go
go build -o bin/golanger cmd/golanger/main.go
go build -o bin/protoc-gen-golangergin cmd/protoc-gen-golangergin/main.go
go build -o bin/template_validator cmd/template_validator/main.go
go build -o bin/impler cmd/impler/main.go