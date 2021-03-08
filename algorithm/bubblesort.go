package main

import "fmt"

func bubblesort(arr []int) {

	sorted := true
	for i:=0; i < len(arr); i++ {

		if !sorted {
			fmt.Println("sorted")
			break
		}

		for j:=0; j < len(arr)-1-i; j++ {

			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
				sorted = true
			}else {
				sorted = false
			}

		}
		fmt.Println(arr)
	}
}

func main() {

	arr := []int{6,1,2,3,4,5}
	bubblesort(arr)
}
