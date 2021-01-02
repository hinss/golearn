package main

import "fmt"

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("Extending slice")
	fmt.Println("arr = ", arr)
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println("s1=%v, len(s1)=%d, cap(s1)=%d", s1, len(s1), cap(s1))
	fmt.Println("s2=%v, len(s2)=%d, cap(s2)=%d", s2, len(s2), cap(s2))

	// 往Slice中添加数据
	// 当超过arr原来的数组的cap时，系统会生成一个新的arr去接
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s2)
	fmt.Println(s3)
	// s4 s5 is no longer view arr
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(arr)
}
