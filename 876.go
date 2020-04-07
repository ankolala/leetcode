package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	cur := head
	size := 1
	for cur.Next != nil {
		size += 1
		cur = cur.Next
	}

	fmt.Println(size)

	mid := size/2
	now := 0
	cur = head
	for head.Next != nil {
		now += 1
		cur = cur.Next
		if now == mid {
			return cur
		}
	}
	return cur
}

func main() {
	fourth := new(ListNode)
	fourth.Val = 6
	fourth.Next = nil

	third := new(ListNode)
	third.Val = 5
	third.Next = fourth

	second := new(ListNode)
	second.Val = 4
	second.Next = third

	first := new(ListNode)
	first.Val = 3
	first.Next = second

	fmt.Println(middleNode(first))

	a:= int64(01|00)
	fmt.Println(a)
}