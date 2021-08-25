package main

import (
	"fmt"
	"learngo/rpc_test/rpc01"
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

//var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(rpc01.HelloServiceName+".Hello", request, reply)
}

/**
实现异步发送rpc请求
本质上就是将 client.Go返回的channel，在需要异步处理的任务后再监听channel中的返回
*/
func doClientWork(client *rpc.Client) {
	helloCall := client.Go(rpc01.HelloServiceName+".Hello", "raymond888", new(string), nil)

	// do some thing
	fmt.Println("do do do do do do do do ")

	helloCall = <-helloCall.Done
	if err := helloCall.Error; err != nil {
		log.Fatal(err)
	}

	args := helloCall.Args.(string)
	reply := helloCall.Reply.(*string)
	fmt.Println(args, *reply)
}

func main() {

	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	doClientWork(client.Client)
	//var reply string
	//err = client.Hello("hello", &reply)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(reply)

}
