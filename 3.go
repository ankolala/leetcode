package main

import (
	"fmt"
	"strings"
)

// 滑动窗口解法
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2{
		return len(s)
	}

	left, right := 0, 0
	max := 1

	for i:=1;i<len(s);i++ {
		if strings.Contains(s[left:right+1], string(s[i])){
			index := left+strings.Index(s[left:right+1],string(s[i]))
			right ++
			left = index +1
			continue
		}
		right ++
		if right-left+1 > max {
			max = right - left + 1
		}

	}

	return max
}

func main() {
	in_str := "abcabcbb"
	in_str = "dvdf"
	in_str = "bbbbb"
	//in_str = "pwwkew"
	//in_str = "a"
	fmt.Println("max size:",lengthOfLongestSubstring(in_str))
}
