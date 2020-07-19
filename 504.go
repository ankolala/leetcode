package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {

	res :=  ""
	if num == 0 {
		return "0"
	}

	// 判断负数
	flag := false
	if num < 0 {
		num = -num
		flag = true
	}

	for num > 0 {
		out1 := num %7
		res = strconv.Itoa(out1) + res

		num = num/7
	}

	if flag {
		res = "-" + res
	}
	return res
}

func main() {

	num := -7

    fmt.Println(num%7, num/7)
	fmt.Println(convertToBase7(num))

}
