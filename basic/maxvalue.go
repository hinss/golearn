package main

// 求最大值的函数
func max(vals ...int) (int, int) {

	maxindex := -1
	max := -1
	for index, val := range vals {
		if val > max {
			maxindex, max = index, val
		}
	}
	return maxindex, max
}

func main() {

}
