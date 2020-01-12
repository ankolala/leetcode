package main

import (
	"fmt"
)

var result = 0

func climbStairs(n int) int {
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
}


func main() {
	fmt.Println(climbStairs(4))
}
