package al

func InsertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {

		j := i
		for arr[j] < arr[j-1] {

			temp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = temp

			j--
			if j == 0 {
				break
			}
		}
	}
	return arr

}
