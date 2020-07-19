package main

import "fmt"

func subarraySum(nums []int, k int) int {
	res , m , sum := 0, make(map[int]int,0),0
	m[0] = 1
	for i:=0;i<len(nums);i++{
		sum += nums[i]
		fmt.Println("sum", sum)
		if _,ok := m[sum-k];ok{
			res += m[sum-k]
		}
		m[sum]++
	}
	return res
}


func main() {
	input := []int{5,4,3,2,4,2}
	fmt.Println(subarraySum(input, 7))
}
