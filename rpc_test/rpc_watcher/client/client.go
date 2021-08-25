package main

import (
	"fmt"
	"learngo/rpc_test/rpc_watcher"
	"log"
	"net/rpc"
	"time"
)

type KVStoreServiceClient struct {
	*rpc.Client
}

func DialKVStoreServiceClient(network, address string) (*KVStoreServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &KVStoreServiceClient{Client: c}, nil
}

func (K KVStoreServiceClient) Get(s string, s2 *string) error {
	return K.Client.Call(rpc_watcher.SERVICE_NAME+".Get", s, s2)
}

func (K KVStoreServiceClient) Set(strings [2]string, s *struct{}) error {
	return K.Client.Call(rpc_watcher.SERVICE_NAME+".Set", strings, s)
}

func (K KVStoreServiceClient) Watch(i int, s *string) error {
	return K.Client.Call(rpc_watcher.SERVICE_NAME+".Watch", i, s)
}

func main() {

	// 建立连接
	client, err := DialKVStoreServiceClient("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 开启一个go routine监听server端的状态
	go func() {

		var keyChanged string
		err := client.Watch(30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("watch:%s has changed!", keyChanged)

	}()

	time.Sleep(3 * time.Second)

	// 休眠3秒之后改变值
	strings := [2]string{"username", "raymond688"}
	err = client.Set(strings, &struct{}{})
	if err != nil {
		log.Fatal(err)
	}

}
