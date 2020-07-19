package main

import "fmt"

/*
给定任意一个字符串str=“…abcxxxdfdllyx…”，请编写一个方法func(str)，经过这个方法处理的str不再包含驼峰字符串。
附：驼峰字符串定义，由3个字符组成，第1个和第3个字符相同且不同于第2个，例如dfd。
 */


 func procStr(input string) (string, bool) {
	 flag := false

 	// 长度小于2直接返回
 	if len(input) <= 2 {
 		return input, flag
	}

	result := ""


	i, j := 0, 2

	for k:=j;k<len(input);k++ {
		// 判断input[i:k]是驼峰字符串
		if string(input[k]) == string(input[i]) && string(input[k]) != string(input[i+1]){
			//fmt.Println(k, string(input[i]), string(input[k]))
			i += 3
			k += 2
			flag = true
			continue
		}

		result = result + string(input[i])
		i += 1
	}


	//处理最后2个字符
	if i+1 < len(input) {
		result = result + string(input[i])
		result = result + string(input[i+1])
	}

	fmt.Println(result, flag)
	if flag == false {
		return result, flag
	} else {
		return procStr(result)
	}
 }




func main() {
	str := "abcylxdfdlcyx"
// abclxllyx-> abclyx
	result, flag := procStr(str)
	fmt.Println(result, flag)
}
