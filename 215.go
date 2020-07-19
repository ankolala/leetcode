package main

import (
	"fmt"
)

func quickSort(input []int, low int, high int) {
	if low >= high {
		return
	}

	i, j := low, high
	pivot := input[i]

	for i < j {
		for input[j] >= pivot && i<j {
			j--
		}

		for input[i] <= pivot && i<j{
			i++
		}

		input[i], input[j] = input[j], input[i]
	}

	input[low] = input[i]
	input[i] = pivot

	quickSort(input, low, i-1)
	quickSort(input, i+1, high)

	return
}

func sortArray(input []int) []int{
	quickSort(input, 0, len(input)-1)
	return input
}

func proc(input []int, k int) int {
	update := []int{}
	hashMap := map[int]int{}

	for i:=0;i<len(input);i++ {
		if _, ok := hashMap[input[i]];!ok {
			update = append(update, input[i])
			hashMap[input[i]] = 1
		}
	}

	sortArray(update)

	return update[k-1]
}


func main() {
	input := []int{1,1,2,3,4,2,3,5}
	fmt.Println(proc(input, 4))
}