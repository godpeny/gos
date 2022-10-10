//go:build tools
// +build tools

// the purpose of this code is to use a tool dependency
// to track the versions of the following executable packages.
// https://github.com/grpc-ecosystem/grpc-gateway#compile-from-source
package tools

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
