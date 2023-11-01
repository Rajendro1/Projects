package main

import (
	grpcserver "main.go/grpcServer"
	httpgin "main.go/httpGin"
)

// protoc --go_out=. --go-grpc_out=. user.proto
func main() {
	go grpcserver.StartGRPCServer()
	httpgin.StartServer()
}