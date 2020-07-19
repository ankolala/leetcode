package main

import (
	"fmt"
)

var result = 0

// 递归算法超出时间限制
func climbStairs(n int) int {
	// 1阶或2阶

	if n == 1 || n == 2 {
		return n
	}

	//a, b := 1, 2

	return climbStairs(n-1) + climbStairs(n-2)
}

// 斐波那契
func climbStairs2(n int) int {

	switch n {
	case 1:
		return 1
	case 2:
		return 2
	default:
		dp1, dp2 := 1, 2
		for i:=2; i<n;i++{
			dp1, dp2 = dp2, dp1+ dp2
		}
		return dp2
	}
}

/*func climbStairs(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	default:
		dp1, dp2 := 1, 2
		for i := 2; i < n; i++ {
			fmt.Println("before",dp1, dp2)
			dp1, dp2 = dp2, dp1+dp2
			fmt.Println(dp1, dp2)
		}
		return dp2
	}
}*/


func main() {
	fmt.Println(climbStairs(4))
}
