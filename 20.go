package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	LeftVec := []string{}

	for i:=0;i<len(s);i++{
		switch string(s[i]) {
		case "(":
			LeftVec = append(LeftVec, ")")
			continue
		case "[":
			LeftVec = append(LeftVec, "]")
			continue
		case "{":
			LeftVec = append(LeftVec, "}")
			continue
		}

		fmt.Println(string(s[i]),LeftVec,LeftVec[len(LeftVec)-1])
		if len(LeftVec)>0 && string(s[i]) == LeftVec[len(LeftVec)-1]{
			LeftVec = LeftVec[:len(LeftVec)-1]
		} else {
			return false
		}


	}
	if len(LeftVec) == 0{
		return true
	}
	return false
}

func main() {

	in := "((())"
	fmt.Println(isValid(in))
}
