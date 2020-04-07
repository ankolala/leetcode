package main

import (
	"fmt"
)


type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	/*
	head:1->2->3-> nil
	pre: 2->3-> nil
	cur:1->2->3-> nil

	head: 1-> nil
	temp:3->nil
	pre: 3-> nil
	cur:  2->1-> nil
	 */

	 //pre, cur := head.Next, head
	 //
	 //head.Next = nil
	 ////fmt.Println(pre,cur, head)
	 //for pre != nil {
	 //	fmt.Println(pre,cur, head)
	 //	temp := pre.Next //2,3,nil
	 //	pre.Next = cur
	 //	cur = pre
	 //	pre = temp
	 //
	 //
	 //}
	 //return cur


	 pre, cur := head.Next, head
	 /*
	 pre: 2->3->nil
	 cur, head: 1-> nil
	  */

	 head.Next = nil

	 for pre != nil {
	 	temp := pre.Next //2->3->nil

	 	pre.Next = cur // 2->1-> nil
		cur = pre
	 	pre = temp
	 }

	 return cur



}

func main() {
	third := new(ListNode)
	third.Val = 3
	third.Next = nil

	second := new(ListNode)
	second.Val = 2
	second.Next = third

	head := new(ListNode)
	head.Val = 1
	head.Next = second

	cur := reverseList(head)
	fmt.Println(cur.Val, cur.Next.Val, cur.Next.Next.Val)
}
