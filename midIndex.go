package main

import "fmt"

/*
输入:
nums = [1, 7, 3, 6, 5, 6]
输出: 3
解释:
索引3 (nums[3] = 6) 的左侧数之和(1 + 7 + 3 = 11)，与右侧数之和(5 + 6 = 11)相等。
同时, 3 也是第一个符合要求的中心索引。

reference url:https://leetcode-cn.com/explore/learn/card/array-and-string/198/introduction-to-array/770/

 */
func pivotIndex(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	var left, right int

	mid := len(nums)/2

	for {

		fmt.Println( mid, left, right)

		for i := 0; i<mid; i++ {
			left = left+nums[i]
		}

		for j:=mid; j<len(nums)-1; j ++ {
			right = right+nums[j]
		}

		if left == right {
			break
		}

		if left < right {
			mid = mid +1
		}

		if left > right {
			mid = mid-1
		}
	}
	fmt.Println( mid, left, right)
	return mid

}

func main() {
	in := []int{1,2,1}
	//in = []int{1,2,3}
	fmt.Println(pivotIndexDouble(in))
}

func pivotIndexDouble(data []int) int {
	if len(data) == 0 {
		return -1
	}
	if len(data) == 1 {
		return 0
	}

	hashMap := map[int]int{}
	sum := 0
	for k, v := range data {
		hashMap[k] = sum
		sum += v
	}

	for i := 0; i < len(data); i ++ {
		if (sum - data[i])/ 2 == hashMap[i] && (sum - data[i]) % 2 == 0 {
			return i
		}
	}

	return -1
}