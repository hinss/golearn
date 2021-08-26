package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"learngo/rpc_test/grpc_certification/proto"
	"log"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile(
		"cert/fakeserver.crt", "server.grpc.io",
	)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost:1234",
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	reply, _ := client.Hello(context.Background(), &proto.String{Value: "raymond9997"})

	fmt.Println(reply.GetValue())

}
