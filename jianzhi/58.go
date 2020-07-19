package main

import "fmt"

func reverseLeftWords(s string, n int) string {

	return s[n:]+s[:n]
}

func main() {
	s := "rachel"
	n := 2

	fmt.Println(reverseLeftWords(s, n)) // expect "chelra"
}
