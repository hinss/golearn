package main

import (
	"fmt"
	"learngo/algorithm/al"
)

func main() {
	arr := []int{4, 7, 9, 5, 10, 3, 14, 1, 18, 20, 0}
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//arr = al.SelectionSort2(arr)
	//fmt.Printf("final arr: %v", arr)

	//arr := []int{8, 7, 6, 5, 4, 3, 2, 1}
	//arr = al.SelectionSort2(arr)
	//fmt.Printf("final arr: %v", arr)

	arr = al.InsertionSort(arr)
	fmt.Printf("final arr: %v", arr)
}
