package al

import "fmt"

func quickSort(arr []int) []int {

	minIndex := 0
	var hasSwapped bool
	for i := 0; i < len(arr); i++ {

		hasSwapped = false
		for j := i; j < len(arr)-1; j++ {
			if arr[j+1] < arr[j] {
				minIndex = j + 1
				hasSwapped = true
			}
		}

		if !hasSwapped {
			return arr
		}

		// swap
		temp := arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = temp

		fmt.Println(arr)
	}

	return arr
}

func main() {

	arr := []int{6, 5, 4, 3, 2, 1}
	arr = quickSort(arr)
	fmt.Printf("final arr: %v", arr)
}
