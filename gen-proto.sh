#!/bin/bash

: ${GOPATH:?"ERROR: GOPATH not set"}

export PATH=$PATH:$GOPATH/bin

protoc_path=$(which protoc)
if [ ! -x "$protoc_path" ]; then
	echo "ERROR: Protobuf Compiler not found"
	exit 1
fi

protoc_gen_go_path=$(which protoc-gen-go)
if [ ! -x "$protoc_gen_go_path" ]; then
	echo "ERROR: Protobuf Go bindings not found"
	echo "go get -u github.com/golang/protobuf/{proto,protoc-gen-go}"
	exit 1
fi

if [ ! -d "protobuf/schema" ]; then
	echo "ERROR: Protobuf Schema not found"
	exit 1
fi

protoc --proto_path=protobuf/schema \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:protobuf \
--go_out=plugins=grpc:protobuf protobuf/schema/university/*.proto