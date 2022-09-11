
FROM golang:1.16 as go_builder

ADD bin /go/bin

# install goreportcard + all relative plugins
RUN git clone --depth 1 https://github.com/gojp/goreportcard.git \
    && cd goreportcard \
    && go install ./vendor/github.com/alecthomas/gometalinter \
    && go install ./vendor/golang.org/x/lint/golint \  
    && go install ./vendor/github.com/fzipp/gocyclo \
    && go install ./vendor/github.com/gordonklaus/ineffassign \
    && go install ./vendor/github.com/client9/misspell/cmd/misspell \
    && go install ./cmd/goreportcard-cli \
    && cd .. && rm -rf goreportcard 

FROM mysql:8 as mysql

FROM golang:1.16
LABEL name="golanger" \
    maintainer="jiandahao" \
    version="v1.0.0" \
    description="Golang, Graphviz, Mysql in a container"

ADD scripts/go_build.sh /usr/local/bin/go_build.sh

# pprof relative
RUN apt-get update && apt-get install graphviz graphviz-doc -y

# goreportcard relative
COPY --from=go_builder /go/bin /go/bin

RUN go get github.com/googleapis/googleapis \
    && go get github.com/grpc-ecosystem/grpc-gateway \
    && go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest \
    && go install github.com/golang/mock/mockgen@v1.6.0
    
# setup mysql env
COPY --from=mysql /usr/sbin/mysqld /usr/sbin/mysqld
COPY --from=mysql /usr/bin/mysql* /usr/bin/
COPY --from=mysql /usr/lib/mysql/private /usr/lib/mysql/private
COPY --from=mysql /usr/lib/x86_64-linux-gnu/libnuma.so.1 /usr/lib/x86_64-linux-gnu/libnuma.so.1
COPY --from=mysql /usr/lib/x86_64-linux-gnu/libaio.so.1 /usr/lib/x86_64-linux-gnu/libaio.so.1
RUN mkdir /var/run/mysqld && mkdir /usr/share/mysql-8.0/ && mkdir /var/lib/mysql && mkdir /var/lib/mysql-files

ENV TZ=Asia/Shanghai
CMD [ "bash","-c", "/usr/local/bin/go_build.sh" ]