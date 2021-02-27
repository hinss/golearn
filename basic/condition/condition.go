package main

import "fmt"

func whetherNumString(s string) bool {
	return s == "3"
}

func main() {

	// 这种 if .. ok的写法 赋值语句的必须是返回bool类型的
	if ok := whetherNumString("3"); ok {
		fmt.Print("ok")
	} else {
		fmt.Print("error")
	}
}
