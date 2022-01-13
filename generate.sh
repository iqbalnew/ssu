#!/bin/bash

swagger="./www/swagger.json"

if [ -f "$swagger" ] ; then
    rm "$swagger"
fi

protoc --proto_path=./proto ./proto/*.proto \
    --proto_path=./vendor \
    --plugin=$(go env GOPATH)/bin/protoc-gen-go \
    --plugin=$(go env GOPATH)/bin/protoc-gen-govalidators \
    --go_out=./server/pb --go_opt paths=source_relative \
    --govalidators_out=./server

protoc --proto_path=./proto ./proto/base.proto \
    --proto_path=./vendor \
    --plugin=$(go env GOPATH)/bin/protoc-gen-grpc-gateway \
    --plugin=$(go env GOPATH)/bin/protoc-gen-openapiv2 \
    --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc \
    --go-grpc_out=./server/pb --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./server/pb \
    --grpc-gateway_opt allow_delete_body=true,logtostderr=true,paths=source_relative,repeated_path_param_separator=ssv \
    --openapiv2_out ./proto \
    --openapiv2_opt logtostderr=true,repeated_path_param_separator=ssv

mv ./proto/base.swagger.json ./www/swagger.json

protoc --proto_path=./proto ./proto/gorm_db.proto \
    --proto_path=./vendor \
    --plugin=$(go env GOPATH)/bin/protoc-gen-gorm \
    --gorm_out=./server