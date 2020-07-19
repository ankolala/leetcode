package main

import (
	"fmt"
)



func maxProfit1(prices []int) int {

	// 零利润
	if len(prices) <= 1 {
		return 0
	}

	/*if len(prices) == 2{
		if prices[1] - prices[0] < 0 {
			return 0
		}
		return prices[1] - prices[0]
	}*/

	var maxProfit, curProfit int

	for i:= 1; i<len(prices); i++{
		for j := 0; j<i; j++ {
			//fmt.Println(i,j)
			curProfit = prices[i] - prices[j]
			if curProfit > maxProfit{
				maxProfit = curProfit
			}
		}
	}
	return maxProfit

}

func maxProfit(prices []int) int {

	var max, cur int

	for i:=1;i<len(prices);i++{
		for j:=0;j<i;j++{
			cur = prices[i] - prices[j]

			if cur > max {
				max = cur
			}
		}
	}
	return max

}


func main() {
	input := []int {7,1,5,3,6,4}
	//input = []int {7,6,4,3,4}
	//input = []int {1,2}
	fmt.Println("max profit:",maxProfit(input))
}
