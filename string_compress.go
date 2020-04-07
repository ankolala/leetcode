package main

import (
	"fmt"
	"strconv"
)

func compressString(S string) string {
	if len(S) == 0 {
		return S
	}

	pivot := S[0]
	result := ""
	cnt := 1
	for i:=1;i<len(S);i++ {
		if S[i] == pivot {
			cnt ++
		} else {
			result += string(pivot)
			result += strconv.Itoa(cnt)
			pivot = S[i]
			cnt = 1
		}
	}
	result += string(pivot)
	result += strconv.Itoa(cnt)
	if len(result) < len(S) {
		return result
	}
	return S
}

func main() {
	in := "abbccd"

	fmt.Println(compressString(in))
}