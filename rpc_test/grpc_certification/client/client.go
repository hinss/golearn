package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learngo/rpc_test/grpc_certification/proto"
	"log"
)

func main() {

	//creds, err := credentials.NewClientTLSFromFile(
	//	"cert/fakeserver.crt", "server.grpc.io",
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//conn, err := grpc.Dial("localhost:1234",
	//	grpc.WithTransportCredentials(creds),
	//)
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &proto.String{Value: "raymond68888"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())

}
