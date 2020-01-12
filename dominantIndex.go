package main

import "fmt"

/*

在一个给定的数组nums中，总是存在一个最大元素 。

查找数组中的最大元素是否至少是数组中每个其他数字的两倍。

如果是，则返回最大元素的索引，否则返回-1。

示例 1:

输入: nums = [3, 6, 1, 0]
输出: 1
解释: 6是最大的整数, 对于数组中的其他整数,
6大于数组中其他元素的两倍。6的索引是1, 所以我们返回1.

reference url:https://leetcode-cn.com/explore/learn/card/array-and-string/198/introduction-to-array/771/


 */

func dominantIndex(nums []int) int {
	if len(nums) == 0{
		return -1
	}

	if len(nums) == 1{
		return 0
	}

	max := 0
	hashMap := map[int]int{}

	for i:=0;i<len(nums);i++{
		hashMap[nums[i]] = i
		if nums[i] > max {
			max = nums[i]
		}
	}

	fmt.Println(max)

	for j:=0;j<len(nums);j++ {
		if nums[j]!=0 && max/nums[j]<2 && j!=hashMap[max]{
			return -1
		}
	}
	return hashMap[max]
}

func main() {
	in := []int{0,0}
	fmt.Println(dominantIndex(in))
}
