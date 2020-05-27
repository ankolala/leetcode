package main

import "fmt"

func lastRemaining(n int, m int) int {
	array := make([]int,n)
	for i:=0; i<n;i++ {
		array = append(array, i)
	}

	result := make([]int, n)
	for len(result) != 1 {

		j := 0
		for ele := range array {
			if j != m - 1  {
				result = append(result, ele)
			}
			j ++
		}
	}
	return result[0]
}

func main() {
	fmt.Println(lastRemaining(5,3))
}
