package exercise

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

func isPalindrome2(x int) bool {

	s := strconv.Itoa(x)
	re := reverseString(s)
	runes := []rune(s)
	reRunes := []rune(re)
	for i, v := range runes {

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

/**
  初始	12321   0
   1    1232	1
   2    123     12
   3    12      123
*/
func isPalindrome(x int) bool {

	if x <= 0 || x%10 == 0 {
		return false
	}

	reverse := 0
	for x > reverse {
		v := x % 10
		reverse = reverse*10 + v
		x = x / 10
	}

	return x == reverse || x == reverse/10
}

func TestLeetcode09() {

	var x int
	x = 1232122
	fmt.Printf("%v: %v\n", x, isPalindrome(x))

	x = 10
	fmt.Printf("%v : %v\n", x, isPalindrome(x))

	x = -121
	fmt.Printf("%v : %v\n", x, isPalindrome(x))

	x = 12321
	fmt.Printf("%v : %v\n", x, isPalindrome(x))
}
