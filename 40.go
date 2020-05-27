package main

import (
	"fmt"
	"sort"
)

func quickSort(arr []int, low int, high int) {
	if low >= high {
		return
	}

	i, j := low, high
	pivot := arr[i]

	for i<j {
		for arr[j]>=pivot&&i<j{
			j--
		}

		for arr[i]<=pivot&&i<j{
			i++
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[low] = arr[i]
	arr[i] = pivot

	quickSort(arr, low, i-1)
	quickSort(arr, i+1, high)
}


func getLeastNumbers(arr []int, k int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr[:k]

}


func ConvIDLongToShort(longID uint64) (shortID uint64) {
	shortID = longID & 0xffffffffffff
	return
}

func main() {
	arr := []int{1,1,4,7,5,4}
	sort.Ints(arr)
	fmt.Println(arr)
	k := 3
	fmt.Println(getLeastNumbers(arr, k))
	fmt.Println(ConvIDLongToShort(17594616345752))
}
