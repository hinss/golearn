package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"learngo/rpc_test/grpc_stream/proto"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	stream, err := client.Channel(context.Background())
	if err != err {
		log.Fatal(err)
	}

	count := 0
	go func() {
		for {
			if err := stream.Send(&proto.String{Value: fmt.Sprintf("Raymond%d 8888", count)}); err != nil {
				log.Fatal(err)
			}

			count++
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}

}
