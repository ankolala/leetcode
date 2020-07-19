package main

import "fmt"

//48, 64, 97
// 97-65=32
func isPalindrome(s string) bool {
	array := []int32{}
	for _, elem := range s{
		if  (47 < elem) && (elem < 58) || (64 < elem) && (elem < 91) || (96 < elem) && (elem < 123) {
			if (64 < elem) && (elem < 91) {
				array = append(array, elem+32)
			} else {
				array = append(array, elem)
			}

		}
	}
	fmt.Println(len(array), array)

	i, j := 0, len(array) -1
	// array长度为偶数
	if len(array) % 2 == 0 {

		for i < j {
			if array[i] != array[j] {
				return false
			}
			i ++
			j --
		}

	} else { // array长度为奇数
		for i != j {
			fmt.Println(i, j)
			if array[i] != array[j] {
				return false
			}
			i ++
			j --
		}
	}
	return true
}

func main() {
	str := "A man, a plan, a canal: Panama"
	str = "race a car"
	fmt.Println(isPalindrome(str))
}

// 65+25= 100
//9