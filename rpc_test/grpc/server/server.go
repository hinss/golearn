package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"learngo/rpc_test/grpc/proto"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *proto.String) (*proto.String, error) {
	reply := &proto.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	log.Println("GRPC Server Start!")
	grpcServer := grpc.NewServer()

	proto.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	// 注册反射服务，便于grpcurl来发送请求测试server的接口
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", ":10001")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
