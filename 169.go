package main

import (
	"fmt"
	"sort"
)


func majorityElement(nums []int) int {

	arraySize := len(nums)
	hashMap := map[int]int{}
	max := 0

	// 数组为空返回-1
	if arraySize == 0 {
		return -1
	}

	// 数组长度为1返回第一个元素
	if arraySize == 1 {
		return nums[0]
	}

	// 遍历元素并存到map中
	for _, elem := range nums {
		if _, ok := hashMap[elem]; ok {
			hashMap[elem] += 1
		} else {
			hashMap[elem] = 1
		}

		if hashMap[elem] > max {
			max = hashMap[elem]
			if max > arraySize/2 {
				return elem
			}
		}
	}


	return -1
}

func main() {
	nums := []int{2,3,1,1,1,2,2}
	sort.Ints(nums)
	fmt.Println(nums)

	fmt.Println(majorityElement(nums))
}
