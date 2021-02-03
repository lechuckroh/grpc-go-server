#!/bin/bash

function generate() {
  local PROTO_PATH=$1
  local PROTO_FILES=$2
  local OUTPUT_DIR=$3

  echo Generating "$PROTO_FILES"

  mkdir -p "$PROTO_PATH/output"

  docker run --rm \
    --entrypoint protoc \
    -v "$PROTO_PATH":/defs \
    --user $UID \
    namely/protoc-all:1.33_0 \
    --proto_path=/defs \
    --proto_path=/opt/include \
    --go_out=plugins=grpc:"$OUTPUT_DIR" \
    --grpc-gateway_out=logtostderr=true,register_func_suffix=Gw:"$OUTPUT_DIR" \
    "$PROTO_FILES"
}

CWD=$(pwd)
OUTPUT_DIR_NAME=output
OUTPUT_PATH=$CWD/proto/$OUTPUT_DIR_NAME

generate "$CWD/proto" hello/hello.proto $OUTPUT_DIR_NAME
cp -Rf "$OUTPUT_PATH/" "$CWD"
rm -rf "$OUTPUT_PATH"
