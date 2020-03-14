package main

import "fmt"

func reverseString(s []byte)  {
    // 双指针法
	left := 0
	right := len(s)-1

	for right > left {
		tmp := s[right]
		s[right] = s[left]
		s[left] = tmp
		left ++
		right --
	}
	return
}

func main() {
	in := "hello"
	tmp := []byte(in)

	reverseString(tmp)
	for _, item := range tmp{
		fmt.Println(string(item))
	}
	fmt.Println(tmp)

}
