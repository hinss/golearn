package exercise

func moveZeroes(nums []int) {

	zeroCount := 0
	for i := 0; i < len(nums)-zeroCount; i++ {
		if nums[i] == 0 {
			for j := i; j < len(nums)-zeroCount-1; j++ {
				temp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = temp
			}
			zeroCount++
			i--
		}
	}

}

func TestLettcodeoffer283(nums []int) {

	moveZeroes(nums)
}
