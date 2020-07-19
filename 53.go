package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

// 贪心算法 {-2,1,-3,4,-1,2,1,-5,4}
func maxSubArray(nums []int)int{
	max, sum := 0, 0
	for i:=0;i<len(nums);i++{
		if nums[i] + sum < 0 {
			sum = 0
		} else {
			sum += nums[i]
		}

		if max < sum {
			max = sum
		}
	}
	return max
}


// 动态规划
func maxSubArray1(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	if len(nums) == 1{
		return nums[0]
	}

	//动态规划
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]

	for i:=1;i<len(nums);i++{
		dp[i] = int(math.Max(float64(dp[i-1]+nums[i]), float64(nums[i])))
		max = int(math.Max(float64(dp[i]), float64(max)))
	}
	return max
}

func main() {
	input := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println("max sub-array:",maxSubArray(input))
}
