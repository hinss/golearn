package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func func1(a, b int, op string) (int, error) {

	switch op {

	case "+":
		return a + b, nil

	case "-":
		return a - b, nil

	case "*":
		return a * b, nil

	case "/":
		return a / b, nil

	default:
		return 0, fmt.Errorf("not support operation: %s", op)

	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

// 函数式编程，函数的参数还是函数
func apply(op func(a, b int) int, a, b int) int {
	// 获得函数的指针
	pointer := reflect.ValueOf(op).Pointer()
	// 通过指针获得函数名
	opName := runtime.FuncForPC(pointer).Name()
	// 打印函数名和传入参数
	fmt.Printf("Calling %s with %d, %d\n",
		opName, a, b)
	return op(a, b)
}

// 参数为可变参数列表的函数
func sumArgs(values ...int) int {
	r := 0
	for i := range values {
		r = r + i
	}
	return r
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {

	//fmt.Println(func1(1,2,"+"))

	// 匿名函数不需要起函数名
	apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4)

	// 求和
	fmt.Println(sumArgs(1, 2, 3, 4, 5))

	// 通过指针交换数据
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	// 常规的交换
	c, d := 4, 5
	c, d = swap2(c, d)
	fmt.Println(c, d)

}
