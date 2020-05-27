package main

import "fmt"

func findRepeatNumber(nums []int) int {
	hashMap := map[int]int{}

	for i:=0;i<len(nums);i++{
		if _,ok := hashMap[nums[i]];ok {
			return nums[i]
		} else {
			hashMap[nums[i]] = 1
		}
	}

	return 0
}

func main() {
	nums := []int{2,2,3}
	fmt.Println(findRepeatNumber(nums))
}
