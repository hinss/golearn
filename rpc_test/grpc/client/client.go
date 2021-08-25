package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"learngo/rpc_test/grpc/proto"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &proto.String{Value: "sbsbsbsb"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
