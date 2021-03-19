package main

import (
	"fmt"
	"learngo/leetcode/exercise"
)

/**
for test
*/
func main() {

	testLeetcodeoffer40Arr := []int{4, 5, 1, 6, 2, 7, 3, 8}
	result := exercise.TestLeetcodeoffer40(testLeetcodeoffer40Arr, 3)

	fmt.Printf("leetcode40 : %v", result)

	testLeetcodeoffer283Arr := []int{0, 1, 0, 3, 12}
	exercise.TestLettcodeoffer283(testLeetcodeoffer283Arr)
	fmt.Printf("leetcode283 : %v", testLeetcodeoffer283Arr)
}
