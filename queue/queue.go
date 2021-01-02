package queue

// 【别名的方式】拓展已存在的切片类型
// 定义一个切片类型实现的Queue
type Queue []int

func (q *Queue) Push(v int) {
	// 必须把append后新的切片赋值给原来的 Queue地址
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	// 注意这里拿出指针切片第一个元素的写法(*q)[0]
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
