package main

import "fmt"

//func permutation(s string) []string {
//	// 存放最终结果
//	var result []string
//
//	// 去重使用
//	hashMap := map[string]int{}
//
//
//	for elem := range s {
//		fmt.Println(string(elem))
//	}
//	return result
//}

func permutation(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	dict := map[string]bool{}
	str := []byte(s)
	fmt.Println(str)
	var f func(index int)
	f = func(index int) {
		if index == len(str) {
			dict[string(str)] = true
			return
		}
		for i := index; i < len(str); i++ {
			tmp := str[index]
			if i == 0 {
				fmt.Println("tmp",tmp)
			}

			str[index] = str[i]
			str[i] = tmp
			f(index + 1)
			str[i] = str[index]
			str[index] = tmp
		}
	}
	f(0)
	res := []string{}
	for k, _ := range dict {
		res = append(res, k)
	}
	return res
}

func main() {
	fmt.Println(permutation("abc"))
}
