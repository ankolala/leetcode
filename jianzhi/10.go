package main

import "fmt"

func numWays(n int) int {

	if n == 0 {
		return 1
	}

	if n == 1 || n == 2 {
		return  n
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i:=2; i<=n; i ++  {
		dp[i]  = (dp[i-1] + dp[i-2]) % 1000000007
	}

	return dp[n]

}


func main() {
	fmt.Println(numWays(3))
}
