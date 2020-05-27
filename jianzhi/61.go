package main

import "fmt"


func isStraight(nums []int) bool {
	count := 0 // 统计大小王的个数
	hashMap := map[int]int{} // 判断非0值重复hashmap
	min, max := 14, 0

	for _, v := range nums {
		if v == 0 {
			count ++
			continue
		}

		if _, ok := hashMap[v]; ok {
			return false
		}
		hashMap[v] ++

		if v > max {
			max = v
		}

		if v < min {
			min = v
		}

	}

	if count == 0 {
		return max - min == 4
	}
	return max - min <= 4
}

func main() {
	// 数组无序
	test1 := []int{3,2,4,5,6}
	//test1 = []int{0,0,1,2,5}

	fmt.Println(isStraight(test1))
}