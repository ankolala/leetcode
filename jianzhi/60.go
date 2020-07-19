package main

import (
	"fmt"
	"strconv"
)

func twoSum(n int) []float64 {

	result := make([]float64, 0)
	if n == 0 {
		return result
	}

	// 最大值
	max := n * 6


	one, _ := strconv.ParseFloat(fmt.Sprintf("%.5f",float64(1.0)/float64(6)), 5)

	// 区间[n*1, n*6]
	for i:=n; i<= max; i++ {


		m := groupNum(i)
		rate := float64(1)
		for j:=0;j<m;j++ {
			rate = rate *  one

		}
	}
	return result

}

func groupNum(n int) int {


	return 1
}

func main() {

	// 1-> [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]
	// 2-> [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]
	// 1 <= n <= 11

	fmt.Println(twoSum(3))
}

// 3 -> 1
//