package main

import "fmt"

func sortArrayByParity(A []int) []int {

	after := len(A)-1
	before := 0

	for before < after {
		fmt.Println("before", before, after, A)
		//before 0 2 [7 6 4 1]
		if A[before]%2 > A[after]%2{
			tmp := A[before]
			A[before] = A[after]
			A[after] = tmp
		}
		if A[before]%2 == 0 {
			before ++
		}
		if A[after]%2 !=0 {
			after --
		}

		fmt.Println("after", before, after, A)

	}
	return A
}


func main() {
	A := []int{7,6,4,1}

	fmt.Println(sortArrayByParity(A))
}
