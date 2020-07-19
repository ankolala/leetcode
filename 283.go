package main

import "fmt"

// 两次遍历
func moveZeroes(nums []int)  {

	cur := 0
	for i:=0;i<len(nums);i++{
		if nums[i] != 0 {
			nums[cur] = nums[i]
			cur ++
		}
	}

	for j:=cur;j<len(nums);j++{
		nums[j] = 0
	}

	return

}

func main() {
	nums := []int{1,0,5,0,6} // i=0,j=6;
	moveZeroes(nums)
	fmt.Println(nums)

	test := []int{}
	test = append(test, nums[2:]...)

}

//一次遍历,快慢指针
func moveZeroesNew(nums []int)  {
	for i:=0;i<len(nums);i++{
		for j:=0;j<i;j++{
			if nums[j] == 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
