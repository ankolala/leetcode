package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	for (left < right) {
		mid := (left+right)/2
		if target < nums[mid] {
			right = mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1

}

func main() {

	nums := []int {-1,0,3,4,5,9,12}
	nums = []int {-1}
	target := -1
	fmt.Println(search(nums, target))
}
