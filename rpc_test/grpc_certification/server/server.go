package main

import (
	context "context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"

	"learngo/rpc_test/grpc_certification/proto"
)

type HelloServiceImpl struct {
}

func (h HelloServiceImpl) Hello(ctx context.Context, s *proto.String) (*proto.String, error) {
	return &proto.String{
		Value: "你好:" + s.GetValue(),
	}, nil
}

func main() {

	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterHelloServiceServer(server, HelloServiceImpl{})

	listener, _ := net.Listen("tcp", ":1234")

	server.Serve(listener)
}
