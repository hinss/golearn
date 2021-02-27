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
	file, err := os.OpenFile(
		filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		// err.(*os.PathError) 将err转型为pathError类型 用pathError变量接
		// ok变量可以接受转型的结果
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
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
