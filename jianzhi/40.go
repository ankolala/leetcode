package main

import "fmt"

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

		fmt.Println(arr)
		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[low] = arr[i]
	arr[i] = pivot

	quickSort(arr, low, i-1)
	quickSort(arr, i+1, high)
}

func quickSort1(arr []int, low int, high int) {
	if low >= high {
		return
	}

	i, j := low, high
	qivot := arr[i]

	for i < j {
		for arr[j] >= qivot && i<j {
			j--
		}

		for arr[i] <= qivot && i<j{
			i++
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[low] = arr[i]
	arr[i] = qivot

	quickSort1(arr, low, i-1)
	quickSort1(arr, i+1, high)
}

func main() {
	arr := []int{5,6,7,9,1,3,4,5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}