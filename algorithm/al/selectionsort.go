package al

import "fmt"

func SelectionSort(arr []int) []int {

	times := 0
	var minIndex int
	for i := 0; i < len(arr)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		// swap
		temp := arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = temp

		times++
		fmt.Printf("arr: %v deep: %d \n", arr, times)
	}

	return arr
}

// 二元选择排序法
func SelectionSort2(arr []int) []int {
	times := 0
	var minIndex, maxIndex int
	for i := 0; i < len(arr)/2; i++ {
		minIndex = i
		maxIndex = i

		for j := i + 1; j < len(arr)-i; j++ {
			if arr[minIndex] > arr[j] {
				//记录最小值下标i
				minIndex = j
			}
			if arr[maxIndex] < arr[j] {
				//记录最大值下标
				maxIndex = j
			}
		}

		if minIndex == maxIndex {
			break
		}

		// 交换最小数到minIndex
		var temp int
		temp = arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp

		if maxIndex == i {
			maxIndex = minIndex
		}

		// 交换最大数到minIndex
		lastIndex := len(arr) - i - 1
		temp = arr[lastIndex]
		arr[lastIndex] = arr[maxIndex]
		arr[maxIndex] = temp

		times++

		fmt.Printf("arr: %v deep: %d \n", arr, times)
	}

	return arr
}
