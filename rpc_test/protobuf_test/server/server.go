package main

import (
	"fmt"
	"learngo/rpc_test/protobuf_test/proto"
)

func main() {

	hq := proto.HelloRequest{
		// 1.one of 可以给field设置其中一种的类型
		Name: &proto.HelloRequest_TrueName{TrueName: "raymondTrue"},
		//Name: &proto.HelloRequest_NickName{NickName: "raymondNickName"},
		// 2.设置枚举值
		Type: proto.NameType_NickName,
		// 3.切片
		Moneys: []float32{8888.88, 6666,88},
		// 4.mymap
		Names: map[string]string{
			"name": "raymond",
			"age": "18",
		},
	}

	fmt.Printf("%v", hq)
}


