package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// 列数起，列数止，行数起，行数止
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1

	result := []int{}

	for ;; {
		// 从左到右
		for i:=l;i<=r;i++{
			result = append(result, matrix[t][i])
		}
		t ++
		if t > b {
			break
		}
		// 从上到下
		for i:=t;i<=b;i++ {
			result = append(result,matrix[i][r])
		}
		r --
		if l > r {
			break
		}

		// 从右到左
		for i:=r;i>=l;i--{
			result = append(result, matrix[b][i])
		}
		b --
		if t > b {
			break
		}

		// 从下到上
		for i:=b;i>=t;i-- {
			result = append(result, matrix[i][l])
		}

		l ++
		if l > r {
			break
		}
	}
	return result
}

func main() {

	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	// 预期：[1,2,3,6,9,8,7,4,5]
	fmt.Println(spiralOrder(matrix))
}
