package main

import "fmt"

func quick_sort(arr []int, low int, high int){


	//递归终止条件
	if low >= high {
		return
	}

	i,j := low, high //新增变量i,j
	pivot := arr[i]

	for i<j {
		for arr[j]>=pivot && i<j {
			j --
		}

		for arr[i]<=pivot&& i<j {
			i ++
		}
		arr[i], arr[j] = arr[j], arr[i]

	}

	arr[low] = arr[i]
	arr[i] = pivot

	quick_sort(arr,low, i-1)
	quick_sort(arr, i+1, high)
}


func main() {
	in := []int{1,2,3}

	/*
	{2,3,4,1,0,5}
	{2,5,4,1,3,5}
	{2,5,1,4,3,5}
	{1,5,2,4,3,5} base =2
	 */

	quick_sort(in, 0,2)

	fmt.Println(in)
}

