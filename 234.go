package main

import (

	//"fmt"
	"fmt"
	"time"
)

type ListNode struct {
	Val 	int
	Next 	*ListNode
}

func newNode(head *ListNode) *ListNode {
	var pre, after *ListNode

	cur := head

	for cur != nil {
		after = cur.Next
		cur.Next = pre

		pre = cur
		cur = after
	}

	return pre
}

// 链表反转
// 1->2->3
func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	if newNode(head) == head {
		return true
	}
	fmt.Println(newNode(head), head)



	return false

}

// 链表Val提取
func isPalindrome1(head *ListNode) bool {

	// 定义数组，存放链表元素
	arrayList := make([]int, 0)

	if head == nil || head.Next == nil {
		return true
	}

	// 遍历链表存放Val
	for head != nil {
		arrayList = append(arrayList, head.Val)
	}

	if len(arrayList) %2 == 0 {
		for i:=0;i<len(arrayList);i++{
			if arrayList[i] != arrayList[len(arrayList)-1-i] {
				return false
			}
		}
		return true
	} else {
		for i:=0;i<len(arrayList);i++{
			if arrayList[i] != arrayList[len(arrayList)-1-i] && i != len(arrayList)/2{
				return false
			}
		}
		return true
	}

	return false

}


func main() {

	next := &ListNode{
		Val : 1,
		Next: nil,
	}
	head := &ListNode{
		Val : 2,
		Next: next,
	}
	// {1,2,3,4,5} 5,2
	isPalindrome(head)




}

