package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	panic("aaa")
	defer fmt.Println(2)
	fmt.Println(3)
}

func writeFile(filename string) {
	// 打开资源
	file, err := os.Create(filename)
	if err != nil {
		// 打开你失败直接抛出异常
		panic(err)
	}
	// 声明defer 在方法结束或者后续出现异常前调用关闭资源
	defer file.Close()

	// 用buffer组装资源
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fbn()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

	fmt.Println(f())

}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
