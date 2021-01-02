package main

import (
	"fmt"
	"strings"
)

func findNonRepeatingSubstr(s string) (int, int) {
	// 记录每个字符最新index的map
	lastOccured := make(map[rune]int)
	// 用于表示最长不含有重复子串的开始下标
	start := 0
	// 用于记录最大长度
	maxLength := 0
	// 用于记录最大长度的index
	maxIndex := 0

	for i, ch := range []rune(s) {
		lastI, ok := lastOccured[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i+1-start > maxLength {
			maxLength = i + 1 - start
			maxIndex = i
		}
		lastOccured[ch] = i
	}

	return maxLength, maxIndex
}

func main() {
	s := "ababccehjkl"
	maxLength, maxIndex := findNonRepeatingSubstr(s)
	// 最后的下标
	lastIndex := maxIndex + 1
	// 计算开始的下标
	startIndex := lastIndex - maxLength
	// reslice截取字符串
	substr := s[startIndex:lastIndex]
	fmt.Println(substr)

	fields := strings.Fields(s)
	fmt.Println(fields)
}
