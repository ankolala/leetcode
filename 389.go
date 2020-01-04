package main

import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
	s = s + t
	hashMap := map[string]int{}

	for i := 0; i < len(s); i ++ {
		if _, ok := hashMap[string(s[i])]; ok {
			hashMap[string(s[i])] += 1
		} else {
			hashMap[string(s[i])] = 1
		}
	}

	for k, v := range hashMap {
		if v % 2 == 1 {
			b := []byte(k)
			return b[0]

		}
	}

	return byte(1)

}

func main() {
	s := "a"
	t := "aa"
	fmt.Println(findTheDifference(s, t))
}
