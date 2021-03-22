package al

func InsertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {

		j := i
		for arr[j] < arr[j-1] {

			temp := arr[j]
			arr[j-1] = arr[j]
			arr[j] = temp

			j--
		}
	}
	return arr

}
