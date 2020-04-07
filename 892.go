package main

import "fmt"

func surfaceArea(grid [][]int) int {

	ans := 0
	for i, _ := range grid { //0, {1,2}; 1,{3,4}
		for j, _ := range grid[i] { // {1,2} or {3,4}
			if 	grid[i][j] == 0 {
				continue
			}
			ans += grid[i][j] * 4 + 2 // 先加完6个面的表面积，之后再精细化处理
			if i > 0 {
				ans -= 2 * min(grid[i-1][j], grid[i][j]) //以最小的方块计算，因为之前min的方块没有减去表面积，所以这里*2
			}

			if j > 0 {
				ans -= 2 * min(grid[i][j-1], grid[i][j])
			}

		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {

	grid := [2][2]int{{1,2},{3,4}}
	fmt.Println(surfaceArea(grid))

}

