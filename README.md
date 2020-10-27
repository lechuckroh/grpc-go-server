# gRPC Server

## Quick Start

```bash
# install golang
$ tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz
$ export PATH=$PATH:/usr/local/go/bin

# set GOPATH
$ mkdir $HOME/go
$ export GOPATH=$HOME/go
$ export PATH=$GOPATH/bin:$PATH

# install protobuf packages
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
$ go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
$ go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

# install project dependencies
$ go get
$ go mod vendor

# build binary
$ make build

# run server
$ ./api-service

# HTTP call
$ curl http://localhost:9080/healthcheck
{}
```
