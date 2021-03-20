package al

import "fmt"

func BubbleSort(arr []int) {

	sorted := true
	for i := 0; i < len(arr); i++ {

		if !sorted {
			fmt.Println("sorted")
			break
		}

		for j := 0; j < len(arr)-1-i; j++ {

			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
				sorted = true
			}

		}
	}
}
