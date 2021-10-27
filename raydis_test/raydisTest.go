package main

import (
	"fmt"
	"gitee.com/ChengHoHins/raydis"
)

func main() {

	config := raydis.Config{
		Addr:         "172.16.178.225",
		Port:         "8000",
		ProtocolType: "tcp",
	}

	client := raydis.NewClient(config)
	err := client.Set("name", "rayych")
	if err != nil {
		panic(err)
	}

	res, err := client.Get("name")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
