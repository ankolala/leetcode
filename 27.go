package main

import "fmt"

func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	i := 0
	j := 0
	for j < length {
		if nums[j] == val {
			// 去找一个不是 val 的值
			j++
		} else {
			// 赋值
			nums[i] = nums[j]
			i++
			j++
		}
	}

	return length - (j - i)
}

func main() {
	in := []int{0,1,2,2,3,0,4,2}
	fmt.Println(removeElement(in, 2))
}
