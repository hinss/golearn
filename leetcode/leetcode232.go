package main

import "fmt"

/**
  用栈实现队列
*/

type MyQueue struct {
	Stack1 []int
	Stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {

	// pop Stack2 中所有元素 入栈 Stack1
	for len(this.Stack2) > 0 {
		// 后去Stack2栈顶元素
		head := this.Stack2[0]
		this.Stack2 = this.Stack2[1:]

		// 入栈 Stack1
		this.Stack1 = append(this.Stack1, head)
	}

	// 待Push的元素入栈 Stack1
	this.Stack1 = append(this.Stack1, x)

	// pop Stack1 中的元素 入栈 Stack2
	for len(this.Stack1) > 0 {
		// 后去Stack2栈顶元素
		head := this.Stack1[0]
		this.Stack1 = this.Stack1[1:]

		// 入栈 Stack1
		this.Stack2 = append(this.Stack2, head)
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	head := this.Stack2[0]
	this.Stack2 = this.Stack2[1:]
	return head
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	head := this.Stack2[0]
	return head
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.Stack2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {

	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj)
	param_2 := obj.Pop()
	fmt.Println(param_2)
	param_3 := obj.Peek()
	fmt.Println(param_3)
	param_4 := obj.Empty()
	fmt.Println(param_4)
}
