package main

import (
	"fmt"
	"learngo/datastructure/stack"
)

/**
for testing
*/
func main() {

	stack := stack.Construct()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.IsEmpty())
	stack.Pop()
	fmt.Println(stack.IsEmpty())

}
