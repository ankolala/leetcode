
/*
给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]

输出:  [1,2,4,7,5,3,6,8,9]

url: https://leetcode-cn.com/explore/learn/card/array-and-string/199/introduction-to-2d-array/774/

 */
package main

import (
	"fmt"

)

func findDiagonalOrder(matrix [][]int) []int {
	out := []int{}
	max := (len(matrix)-1) + (len(matrix[0])-1) //m+n最大值
	fmt.Println(max)
	for i:=0;i<max;i++{
		out = append(out, matrix[i][0])
	}
	return out
}

func main() {

	matrix := make([][]int,1)
	matrix[0] = []int{1,2} //[0,0][0,1][0,2]
	matrix[1] = []int{4,5} //[1,0][1,1][1,2]
	matrix[2] = []int{7,8} //[2,0][2,1][2,2]
	fmt.Println(matrix)
	fmt.Println(findOrder(matrix))
	//[1,2,4,7,5,3,6,8,9]
}


func findOrder(matrix [][]int) []int {

	x,y := 0, 0
	if len(matrix) == 0 {
		return []int{}
	}

	length, height := len(matrix[0]), len(matrix)
	ans := []int{}

	ans = append(ans, matrix[y][x])

	for x < length - 1 || y < height - 1 {
		//左下方向
		if x < length - 1 {
			x ++
		} else {
			y ++
		}
		for x > 0 && y < height - 1{
			ans = append(ans, matrix[y][x])
			x --
			y ++
		}
		ans = append(ans, matrix[y][x])
		if x >= length - 1 && y >= height -1 {
			break
		}

		//右上方向
		if y < height - 1 {
			y ++
		} else {
			x ++
		}

		for y > 0 && x < length - 1{
			ans = append(ans, matrix[y][x])
			x ++
			y --
		}
		ans = append(ans, matrix[y][x])
	}

	return ans

}
