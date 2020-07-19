package main

import "fmt"

func singleNumber1(nums []int) int {

	hashMap := map[int]int{}

	for i:=0;i<len(nums);i++{
		if _,ok := hashMap[nums[i]];ok {
			hashMap[nums[i]] += 1
		} else {
			hashMap[nums[i]] = 1
		}
	}

	for k, v := range hashMap {
		if v == 1 {
			return k
		}
	}

	return -1
}

func singleNumber(nums []int) int {
	// 任何数和0做异或操作，仍未原值
	// 任何数和自身异或，结果为0
	single := 0

	for _, num := range nums {
		single ^= num
	}

	return single
}

func main() {
	input := []int{4,1,2,1,2}
	fmt.Println(singleNumber(input))
}
