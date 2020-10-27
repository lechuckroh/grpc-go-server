#!/bin/bash
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH

OUTPUT=./

function generate() {
  PROTO_PATH=$1
  PROTO_FILES=$2
  protoc \
    --proto_path=$PROTO_PATH \
    --go_out=plugins=grpc:$OUTPUT \
    --grpc-gateway_out=logtostderr=true,register_func_suffix=Gw:$OUTPUT \
    --swagger_out=logtostderr=true:$OUTPUT \
    $PROTO_FILES
}

generate . proto/hello/hello.proto
