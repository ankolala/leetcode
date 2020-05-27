package main

import "fmt"

func maxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	curSum, maxSum := nums[0], nums[0]
	// 注意处理边界，i从1开始
	for i:=1;i<len(nums);i++ {
		if curSum < 0 {
			// 当前累计和小于0，则从下个元素开始计算
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}

		if maxSum < curSum {
			maxSum = curSum
		}
	}
	return maxSum
}

func main() {
	nums := []int{1,2}
	fmt.Println(maxSubArray(nums))
}