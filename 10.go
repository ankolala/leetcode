package main

import "fmt"

//超出时间限制
//func fib(n int) int {
//
//	if n == 0 || n == 1{
//		return n
//	}
//
//	 return fib(n-1) + fib(n-2)
//
//}


func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	a, b, temp := 0, 1, 0

	for i:=2;i<=n;i++ {
		temp = b
		b = (a + b) % 1000000007
		a = temp
	}
	return b
}


func main() {
	fmt.Println(fib(3))
}
