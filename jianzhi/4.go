package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {

	if matrix == nil || len(matrix) == 0|| len(matrix[0]) == 0 {
		return false
	}

	rows := len(matrix)
	row, column := 0, len(matrix[0])-1

	for row<rows && column >= 0 {

		if target == matrix[row][column] {
			return true
		} else if target > matrix[row][column] {
			row ++
		} else if target < matrix[row][column]{
			column --
		}
	}


	return false
}

func main() {

	//matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	matrix := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
	target := 0
	fmt.Println(findNumberIn2DArray(matrix, target))
}
