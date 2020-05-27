package main

import "fmt"

func replaceSpace(s string) string {
	result := ""
	for _,item := range s {
		if string(item) == " " {
			result += "%20"
		} else {
			result += string(item)
		}
	}
	return result
}

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
	//We%20are%20happy.

}
