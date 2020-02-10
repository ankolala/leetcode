package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	out := ""

	if len(s) <= 1 {
		return s
	}

	if strings.Contains(s, " ") == false {
		for i:=len(s)-1;i>=0;i--{
			out = out + string(s[i])
		}
		return out
	} else {
		lastSpace := -1

		for j:=0;j<=len(s)-1;j++{
			if string(s[j]) == " "{
				for i:=j-1;i>lastSpace;i--{
					out = out + string(s[i])
				}
				out = out + " "

				lastSpace = j
			}
		}
		for i:=len(s)-1;i>lastSpace;i--{
			out = out + string(s[i])
		}
	}
	return out
}

func main() {
	in := "Let's take LeetCode contest"
	//in = "let's go"
	fmt.Println(string(in))
	fmt.Println("out:",reverseWords(in))
	fmt.Println(reverseWords(in) == "s'teL ekat edoCteeL tsetnoc" )
}
