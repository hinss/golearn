package main

/**
	以下示例显示了用两个不同的Sum类型方法处理累加
 */

import "fmt"

// SumInts 累加map里的int64类型的value
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats 累加map里的float64类型的value
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// 初始化测试int64类型累加的map
	ints := map[string]int64{
		"first": 34,
		"second": 12,
	}

	// 初始化测试float64类型累加的map
	floats := map[string]float64{
		"first": 35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

}
