package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 示例一: 使用闭包实现斐波那契数列
// 1 1 2 3 5 8 13 21 34 55 89
func fbn() idGen {
	a, b := 0, 1
	return func() int {
		a, b = b, b+a
		return a
	}
}

// 示例二: 为函数实现接口
// 1.首先要定义type类型 是类型才能实现接口
type idGen func() int

// 2.函数作为一等公民也作为接收者
// 2.实现io.Reader接口
func (g idGen) Read(p []byte) (n int, err error) {
	// 2.1 调函数取得下一个元素
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	// 2.2 把元素写进去p []byte数组中
	s := fmt.Sprintf("%d\n", next)
	// 2.3 把字符串写到 字节数组中
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	f := fbn()

	printFileContents(f)

	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

}
