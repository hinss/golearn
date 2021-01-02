package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

// 这里就是使用者方定义接口
type Retriever interface {
	Get(url string) string
}

// 定义使用者的函数 download
func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	// 传指针和值都可以
	//r = &mock.Retriever{"this is a fake imooc.com"}
	r = mock.Retriever{"this is a fake imooc.com"}
	inspect(r)

	r = &real.Retreiver{
		UsersAgent: "Mozilla/5.0",
		TimeOut:    time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetreiver := r.(*real.Retreiver)
	fmt.Println(realRetreiver.UsersAgent)

	//fmt.Println(download(r))

}

// 检查接口的类型
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	// switch 接口的类型
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contens: ", v.Contents)
	case *real.Retreiver:
		fmt.Println("UsersAgent: ", v.UsersAgent)
	}
}
