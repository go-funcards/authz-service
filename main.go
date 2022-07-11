package main

import authz "github.com/go-funcards/authz-service/cmd"

//go:generate protoc -I proto --go_out=./proto/v1 --go-grpc_out=./proto/v1 proto/v1/checker.proto proto/v1/definition.proto proto/v1/rule.proto proto/v1/subject.proto

func main() {
	authz.Execute()
}
