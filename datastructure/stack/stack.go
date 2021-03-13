package stack

/**
使用切片实现一个栈
*/
type Stack struct {
	Values []int
}

func Construct() Stack {
	return Stack{}
}

func (this *Stack) Push(val int) {
	this.Values = append(this.Values, val)
}

func (this *Stack) Pop() int {
	// 返回最后一个元素
	last := this.Values[len(this.Values)-1]
	// 缩容
	this.Values = this.Values[:last-1]

	return last
}

func (this *Stack) IsEmpty() bool {

	return len(this.Values) == 0
}
