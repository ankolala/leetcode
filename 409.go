package main

import "fmt"

func longestPalindrome(s string) int {
	hashMap := map[string]int{}
	for _, item := range s {
		hashMap[string(item)] += 1
	}
	count := 0
	for _, v := range hashMap{
		count += v/2*2 // 比如a出现了5次，其实最终只能保留4次，奇数位单算
		if count%2==0 && v%2 == 1{ //只保留1个奇数位
			count += 1
		}

	}
	return count

}

func main() {
	ss := "aaa"
	fmt.Println(longestPalindrome(ss))
}
