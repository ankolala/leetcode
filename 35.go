package main

import "fmt"

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)
	for (low < high) {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			low = mid + 1
		}

		if nums[mid] > target{
			high = mid
		}
	}
	return high

}

func main() {
	nums := []int {1,3,5}
	num := 5
	fmt.Println("search insert:", searchInsert(nums, num))
}
