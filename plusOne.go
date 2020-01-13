package main

import (
	"fmt"
)

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。

reference url: https://leetcode-cn.com/explore/learn/card/array-and-string/198/introduction-to-array/772/
 */

func plusOne(digits []int) []int {

	if len(digits) == 0 {
		return []int{1}
	}

	out := make([]int, len(digits)+1)
	/*if len(digits) == 1{
		if digits[0] + 1 < 10 {
			return []int{digits[0]+1}
		}
		out[1] = (digits[0]+1)%10
		out[0] = 1
		return out

	}*/

	len := len(digits)-1

	if digits[len] + 1 < 10 {
		digits[len] = digits[len]+1
		return digits
	} else {
		last := (digits[len]+1)%10
		out = plusOne(digits[:len])
		out = append(out, last)
		return out
	}

	return []int{1}
}

func main() {
	in := []int{9,9,9,9,9}
	in = []int{9}
	fmt.Println(plusOne(in))
}
