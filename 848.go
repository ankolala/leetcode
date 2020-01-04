package main

import "fmt"

func shiftingLetters(S string, shifts []int) string {

}

func main() {
	S := "abc"
	shifts := []int{3,5,9}
	fmt.Println(shiftingLetters(S,shifts))

}

/*
1 <= S.length = shifts.length <= 20000
0 <= shifts[i] <= 10 ^ 9

 */