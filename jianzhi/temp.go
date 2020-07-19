package main

import "fmt"

func reverseOnlyLetters(S string) string {
	if len(S) <= 1{
		return S
	}
	//非字母用hashMap存储，字母用数组存储

	hashMap := map[int]string{}
	saveArray := make([]string, 0)
	result := ""

	for i:=0;i<len(S);i++{
		if string(S[i]) == "-"{
			fmt.Println(S[i])
		}
		if (S[i] >= 65 && S[i]<=90) || (S[i] >= 97 &&S[i] <= 122) {
			saveArray = append(saveArray, string(S[i]))
		} else {
			hashMap[i] = string(S[i])
			fmt.Println(i, string(S[i]))
		}
	}

	end := len(saveArray) - 1
	for j:=0;j<len(S);j++{
		// 非字母直接追加
		if _, ok := hashMap[j];ok {
			result = result + hashMap[j]
		} else {
			result = result + saveArray[end]
			end = end - 1
		}
	}
	return result

}

func main() {
	input := "a-bC-dEf-ghIj"
	fmt.Println(reverseOnlyLetters(input))
}
