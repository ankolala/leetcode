package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	// 边界检查
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	out := make([]int, m*n)

	// 1、计算遍历轮数
	var totalRound int
	var outIndex int
	if (m+1)/2 > (n+1)/2 {
		totalRound = (n+1)/2
	} else {
		totalRound = (m+1)/2
	}
	// 2、每轮遍历，遍历顺序：上、右、下、作
	for round:=0;round<totalRound;round++{
		// 3、横向、纵向各去除最后一位元素，保留至下一轮遍历
		minX, minY, maxX, maxY := round, round, n-1-round, m-1-round
		// 上
		for i:=minX;i<=maxX;i++{
			out[outIndex] = matrix[minY][i]
			outIndex++
		}
		fmt.Println(round,out)
		fmt.Println(minY+1,maxY)
		// 右
		for j:=minY+1;j<=maxY;j++{
			out[outIndex] = matrix[j][maxX]
			outIndex++
		}

		// 下
		if maxY-minY >=1{
			for s:=maxX-1;s>=minX;s--{
				out[outIndex] = matrix[maxY][s]
				outIndex ++
			}
		}
		// 左
		if maxX-minX >= 1 {
			for n:= maxY - 1; n > minY; n-- {
				out[outIndex] = matrix[n][minX]
				outIndex ++
			}
		}

	}
	return out


}

func main() {
	in := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	fmt.Println(spiralOrder(in))
}