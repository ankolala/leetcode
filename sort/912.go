package sort

import (
	"fmt"
)

func sortArray11(nums []int) []int {
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


func sortArray(nums []int) []int {
	var (
		lens = len(nums)
		index int
		min int
	)
	for i,v := range nums{
		index = i
		min = v
		for j:=i+1; j<lens; j++ {
			fmt.Println("index",index, "min",v, "nums", nums)
			if nums[j] < min {
				index = j
				min = nums[j]
			}

		}
		//fmt.Println("index",index, "min",v, "nums", nums)
		if index != i {
			fmt.Println("debug",index, i, nums[index], nums[i])
			nums[i],nums[index] = nums[index],nums[i]
		}
	}
	return nums
}

func main() {
	 in := []int{5,1,1,2,0,0}
	 out := sortArray11(in)

	fmt.Println(out)
}
