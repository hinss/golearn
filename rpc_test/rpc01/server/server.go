package main

import (
	myrpc "learngo/rpc_test/rpc01"
	"log"
	"net"
	"net/rpc"
	"time"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {

	time.Sleep(10 * time.Second)
	*reply = "hello:" + request
	return nil
}

func main() {

	myrpc.RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
