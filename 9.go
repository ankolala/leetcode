package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {

	str := strconv.Itoa(x)

	i, j := 0, len(str)-1
	for i<j {
		fmt.Println(str[i], str[j])
		if str[i] == str[j] {
			i ++
			j --
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(1001))

}
