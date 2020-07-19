package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {

	if nums == nil || len(nums) == 0{
		return []int{}
	}

	if k == 1 {
		return nums
	}


	// 滑动左边界
	i := 0

	result := []int{}

	for j:=k-1;j<len(nums);j++{
		// 初始化max值
		max := 0
		fmt.Println(j)
		for x:=i;x<=j;x++{
			fmt.Println("i=",i, x,nums[x])

			if nums[x] > max {
				max = nums[x]
			}
		}

		result = append(result, max)
		i ++

	}

	return result


}

func main() {

	input := []int{7,2,4}

	fmt.Println(maxSlidingWindow(input,2))
}
