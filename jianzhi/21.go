package main

import "fmt"


func exchange(nums []int) []int {

	if nums == nil  || len(nums) <= 1 {
		return nums
	}
	left, right := 0, len(nums) - 1

	// 偶数位于数组的后半部分
	for left < right {

		for nums[right] % 2  == 0 && left < right{
			right --
		}

		for nums[left] % 2 == 1 && left < right{
			left ++
		}

		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		} else {
			return nums
		}

	}

	return nums
}


func main() {
	in := []int{1,2,3,4} // 1,3,2,4;
	fmt.Println(exchange(in))
}

