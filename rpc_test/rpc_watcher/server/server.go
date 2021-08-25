package main

import (
	"learngo/rpc_test/rpc_watcher"
	"log"
	"net"
	"net/rpc"
)

func main() {

	rpc_watcher.RegisterKVStoreService(rpc_watcher.NewKVStoreService())

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
