package main

import (
	"fmt"
)


func isHappy(n int) bool {
	hash := make(map[int]bool)
	for n!= 1 {
		if _,ok := hash[n]; ok {
			return false
		}
		hash[n] = true

		next := 0
		for n != 0 {
			next += (n%10) * (n%10)
			n = n/10
		}
		n = next
	}

	return true

}


func main() {

	fmt.Println(isHappy(19))
}
