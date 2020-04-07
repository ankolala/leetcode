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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0{
		return float64(0)
	}
	var temp []int
	if len(nums1) > len(nums2) {
		temp = make([]int, len(nums1))
		copy(temp, nums1)
		if len(nums1) != 0 {
			for _, num2 := range nums2 {
				temp = append(temp, num2)
			}
		}
	} else {
		temp = make([]int, len(nums2))
		copy(temp, nums2)
		if len(nums2) != 0 {
			for _, num1 := range nums1{
				temp = append(temp, num1)
			}
		}

	}

	quick_sort(temp, 0, len(nums1)+len(nums2)-1)
	var result float64
	fmt.Println("sort",temp)
	mid := (len(temp)-1)/2
	if len(temp)%2==0{
		fmt.Println(mid)
		result = (float64(temp[mid]) + float64(temp[mid+1]))/2

	} else {
		if len(temp) == 1 {
			return float64(temp[0])
		}
		return float64(temp[mid])


	}
	fmt.Println(temp)
	return result
}

func main() {
	nums1 := []int{}
	nums2 := []int{1,2,3,4,5}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
