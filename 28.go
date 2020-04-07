package main

import "fmt"

func strStr(haystack string, needle string) int {

	size := len(needle)
	if size == 0 {
		return 0
	}
	fmt.Println(size,needle)

	if size == len(haystack) {
		if haystack == needle {
			return 0
		}
	}

	for i:=0;i<len(haystack)-1;i++ {
		fmt.Println(string(haystack[i:i+size]),i)
		if i+size-1 > len(haystack)-1 {
			break
		}

		if string(haystack[i:i+size]) == needle{
			return i
		}

	}
	//cannot find
	return -1
}


func main() {

	haystack := "mississippi"

	needle := "pi"
	fmt.Println(strStr(haystack,needle))
}
