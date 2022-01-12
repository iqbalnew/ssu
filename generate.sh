#!/bin/bash

# go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
# go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

# go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
# go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

# go get -u google.golang.org/protobuf/cmd/protoc-gen-go
# go install google.golang.org/protobuf/cmd/protoc-gen-go

# go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --proto_path=./proto ./proto/*.proto \
    --proto_path=./vendor \
    --plugin=$(go env GOPATH)/bin/protoc-gen-grpc-gateway \
    --plugin=$(go env GOPATH)/bin/protoc-gen-openapiv2 \
    --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc \
    --plugin=$(go env GOPATH)/bin/protoc-gen-go \
    --go_out=./server/pb --go_opt paths=source_relative \
    --go-grpc_out=./server/pb --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./server/pb \
    --grpc-gateway_opt allow_delete_body=true,logtostderr=true,paths=source_relative,repeated_path_param_separator=ssv \
    --openapiv2_out ./proto \
    --openapiv2_opt logtostderr=true,repeated_path_param_separator=ssv