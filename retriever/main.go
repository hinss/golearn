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

type Poster interface {
	Post(url string,
		form map[string]string) string
}

// 定义接口组合 包含Retriever 和 Poster的能力
type RetrieverPoster interface {
	Retriever
	Poster
}

// 定义使用者的函数 download
const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

func session(s RetrieverPoster) string {
	// 可以调Get
	//s.Get()
	// 还可以调Post
	return s.Post(url, map[string]string{
		"contents": "aaaaaaaaa",
	})
}

func main() {
	var r Retriever
	// 传指针和值都可以
	//r = &mock.Retriever{"this is a fake imooc.com"}
	retriever := &mock.Retriever{"this is a fake imooc.com"}
	r = retriever
	inspect(r)

	r = &real.Retreiver{
		UsersAgent: "Mozilla/5.0",
		TimeOut:    time.Minute,
	}
	inspect(r)

	// Type assertion
	if mockRetreiver, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetreiver.Contents)
	} else {
		fmt.Println("not mock Retriever")
	}
	//fmt.Println(download(r))

	fmt.Print(session(retriever))

}

// 检查接口的类型
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	// switch 接口的类型
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contens: ", v.Contents)
	case *real.Retreiver:
		fmt.Println("UsersAgent: ", v.UsersAgent)
	}
}
