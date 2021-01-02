package main

import "fmt"

func main() {

	// 1.定义数组
	var array1 [5]int
	array2 := [3]int{1, 2, 3}
	array3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(array1, array2, array3)
	fmt.Println(grid)

	// 2.遍历数组
	// 2.1 遍历数组下标
	for i := range array3 {
		fmt.Println(array3[i])
	}

	// 2.2 遍历数组下标,对应点值(跟python类似)
	for i, v := range array3 {
		fmt.Println(i, v)
	}

	// 2.3 遍历数组 不要下标只要数值的方式（跟Java 的for类似）
	for _, v := range array3 {
		fmt.Println(v)
	}

}
