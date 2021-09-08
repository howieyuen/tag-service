#!/bin/bash

#protoc --go_out=. --go_opt=paths=source_relative \
#    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
#    --proto_path="${GOPATH}"/src --proto_path=. \
#    proto/*.proto

protoc -I"$GOPATH"/src -I. \
  --grpc-gateway_out=logtostderr=true:. \
  ./proto/*.proto
