package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	i,j := 0, len(nums)-1
	for i<j {
		fmt.Println(i,j)
		if (nums[i] + nums[j] < target) {
			i ++
		}else if (nums[i] + nums[j] > target) {
			j --
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return []int{}
}

func main() {
	input := []int{2,7,11,15}

	fmt.Println(twoSum(input, 9))
}
