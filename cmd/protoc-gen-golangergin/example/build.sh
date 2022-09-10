DIR=$(pwd)
cd ..
go build -o $DIR/protoc-gen-golangergin
cd $DIR
protoc -I=${GOPATH}/src/github.com/googleapis/googleapis -I=${GOPATH}/src -I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway --plugin=./protoc-gen-golangergin \
--golangergin_out=withclient=true:. -I. *.proto

rm protoc-gen-golangergin