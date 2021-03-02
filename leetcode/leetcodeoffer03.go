package main

import "fmt"

/**
在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
请找出数组中任意一个重复的数字。
*/

/**
使用map判断是否存在
*/
func findRepeatNumber(nums []int) int {

	dataMap := make(map[int]int)

	for _, v := range nums {

		if _, ok := dataMap[v]; ok {
			return v
		} else {
			dataMap[v] = v
		}
	}

	return 0
}

/**
  原地置换排序法, 在 在长度为n 元素取值范围是 0 - n-1 的数组里，
  如果没有重复数字, 那么正常排序后, 数字i应该就在下标为i的位置。
*/
func findRepeatNumber2(nums []int) []int {

	var temp int

	for i := 0; i < len(nums); i++ {

		for i != nums[i] {
			// 交换
			temp = nums[i]
			nums[i] = nums[temp]
			nums[temp] = temp
		}
	}

	return nums
}

func main() {

	testArr := []int{2, 3, 1, 0, 5, 4}

	testArr = findRepeatNumber2(testArr)

	fmt.Print(testArr)

}
