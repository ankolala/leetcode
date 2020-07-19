package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	result := ""
	// 考虑进位
	extra := 0

	// 保证num1长度大于num2
	exchange(&num1, &num2)

	j := len(num2)-1
	// 外层遍历长度较长的字符串
	for i:=len(num1)-1;i>=0;i-- {
		var item2 int
		item1, _ := strconv.Atoi(string(num1[i]))

		if j>=0 {
			//fmt.Println(j)
			item2, _ = strconv.Atoi(string(num2[j]))
			j --
			temp := (item1 + item2 + extra) % 10
			extra = (item1 + item2  + extra) / 10

			result = strconv.Itoa(temp) + result
		} else {
			fmt.Println("extra", extra)
			temp := (item1  + extra) % 10
			extra = (item1 + extra) / 10

			result = strconv.Itoa(temp) + result
		}

	}

	if extra == 1 {
		result = "1" + result
	}

	return result
}

func exchange(num1 *string, num2 *string) {
	if (len(*num1) > len(*num2)) {
		return
	}
	temp := *num1
	*num1 = *num2
	*num2 = temp
	return
}

func quickSort(low int, high int, array []int) {
	if low >= high {
		return
	}
	i,j := low, high
	pivot := array[i]

	for i<j {
		for i<j && array[j] > pivot {
			j--
		}

		for i<j && array[i] < pivot{
			i ++
		}

		array[i], array[j] = array[j], array[i]
	}
	temp := array[i]
	array[i] = pivot
	pivot = temp

	quickSort(low, i-1, array[:i])
	quickSort(i+1, high, array[i+1:])

}

func main() {
	//A := "99"
	//B := "9"
	//fmt.Println(addStrings(A, B))

	input := []int{3,4,1,6}
	quickSort(0,3, input)
	fmt.Println(input)
}
