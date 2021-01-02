package main

import "fmt"

// 函数式编程实现累加
func adder() func(v int) int {
	sum := 0

	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {

	// 在go语言中，函数也是一等公民。
	// 这里 a:= adder() 拿到的实际是adder() 的返回值,它的返回值也是一个函数，
	// 这个函数入参是(v int) 所以可以a(i)这样调用
	a := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}

}
