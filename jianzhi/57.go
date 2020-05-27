package main

import "fmt"


func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums) - 1
	for i < j {
		if nums[i] + nums[j] < target {
			i ++
		}

		if nums[i] + nums[j] > target {
			j --
		}

		if nums[i] + nums[j] == target {
			break
		}
	}
	return []int{nums[i], nums[j]}
}

func main() {
	// 递增数组
	nums := []int{0,1,2,7,9}
	fmt.Println(twoSum(nums, 9))
}
