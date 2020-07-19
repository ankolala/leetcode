package main

import (
	"fmt"
	"sort"
)

func findDisappearedNumbers(nums []int) []int {

	if len(nums) <= 1 {
		return []int{}
	}

	result := []int{}

	sort.Ints(nums)
	fmt.Println(nums)
	pivot := nums[0]

	for i:=1;i<len(nums);i++ {
		if nums[i]-pivot>1 {
			fmt.Println(nums[i], pivot)
			for j:=pivot+1;j<nums[i];j++ {
				result = append(result, j)
			}
		}
		pivot = nums[i]
	}

	return result

}

func main() {
	input := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDisappearedNumbers(input))
}
