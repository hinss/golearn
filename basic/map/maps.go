package main

import "fmt"

func main() {

	// 初始化map
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m2, m3)

	// 遍历map
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	// 可以多接受一个key判断 key在map中是否存在
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	// golang当key不存在时，仍然会返回空串，不需要判空
	if cause, ok := m["cause"]; ok {
		fmt.Println(cause, ok)
	} else {
		fmt.Printf("key does not exits")
	}

	fmt.Println("Deleting values")
	delete(m, "name")
	fmt.Println(m)

}
