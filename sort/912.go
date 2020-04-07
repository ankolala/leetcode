package main

import (
	"fmt"
)

func sortArray(nums []int) []int {
	size := len(nums)
	for i:=0;i<size;i++{
		for j:=0;j<size-i-1;j++ {
			if nums[j]>nums[j+1]{
				tmp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = tmp
			}
		}

		}
	return nums
}

func main() {
	 in := []int{5,1,1,2,0,0}
	 out := sortArray(in)

	fmt.Println(out)
}
