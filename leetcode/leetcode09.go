package main

import (
	"fmt"
	"strconv"
)

/**
	判断一个整数是否回文
	回文: 123  321  no
         121  121  yes
         -121 121-  no
 */

func isPalindrome(x int) bool {

	s := strconv.Itoa(x)
	re := reverseString(s)
	runes := []rune(s)
	reRunes := []rune(re)
	for i,v := range runes {

		if v != reRunes[i] {
			return false
		}

	}
	return true
}

// 1234567
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}



func main() {

	var x int
	x = 123
	fmt.Printf("123 : %v\n", isPalindrome(x))

	x = 121
	fmt.Printf("121 : %v\n", isPalindrome(x))

	x = -121
	fmt.Printf("-121 : %v\n", isPalindrome(x))

	//s := "abc"
	//
	//fmt.Printf("reverse abc -> %s", reverseString(s))




}
