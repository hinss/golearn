package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"learngo/rpc_test/grpc_stream/proto"
	"log"
	"net"
)

type HelloServiceImpl struct{}

func (h HelloServiceImpl) Hello(ctx context.Context, s *proto.String) (*proto.String, error) {
	return &proto.String{
		Value: "你好: " + s.GetValue(),
	}, nil
}

func (h HelloServiceImpl) Channel(stream proto.HelloService_ChannelServer) error {

	for {

		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &proto.String{
			Value: "你好:" + recv.GetValue(),
		}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {

	server := grpc.NewServer()
	proto.RegisterHelloServiceServer(server, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(listener)
}
