package main

import (
	"fmt"
	"strings"
)

func boldWords(words []string, S string) string {

	// 双指针记录最长匹配字符串

	i, j := 0, len(S)-1


	fmt.Println(i, j)
	return ""
}


// 给定 words = ["ab", "bc"] 和 S = "aabcd"，需要返回 "a<b>abc</b>d"。

func main() {

	words := []string{"ab", "bc"}

	S := "aabcd"
	idx := strings.Index("aabcd","ab")

	println(idx)
	fmt.Println(boldWords(words, S))

}
