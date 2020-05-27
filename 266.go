package main

import "fmt"

func canPermutePalindrome(s string) bool {
	hashMap := map[string]int{}
	count := 0

	for _, elem := range s {
		if _, ok := hashMap[string(elem)];ok {
			hashMap[string(elem)] += 1
		} else {
			hashMap[string(elem)] = 1
		}
	}

	for _, v := range hashMap{
		if v % 2 != 0{
			count += 1
		}
	}

	return count <= 1
}

func main() {
	s := "carerac"
	fmt.Println(canPermutePalindrome(s))
}