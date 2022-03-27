DIR=$(pwd)
cd ..
go build -o $DIR/protoc-gen-golangergin
cd $DIR
protoc -I=${GOPATH}/src/github.com/googleapis/googleapis -I=${GOPATH}/src --plugin=./protoc-gen-golangergin \
--golangergin_out=. -I. *.proto

rm protoc-gen-golangergin