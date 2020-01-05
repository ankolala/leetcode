package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	     Val int
	     Next *ListNode
}
//origin:
//1(cur) -> 2(cur.Next)-> 3(cur.Next.Next)  nil(pre）

//1st:
//pre: 1-> 2 -> 3 -> nil
//cur: 2 -> 3 -> nil
// cur.next
//2(cur)->3(cur.Next)-> nil(cur.Next.Next)-> 1(pre)

//2nd:
//3(cur)->nil(cur.Next)-> 1(cur.Next.Next)-> 2(pre)

//3rd
//nil(cur)->1(cur.Next)-> 2(cur.Next.Next)-> 3(pre)


func reverseList2(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	fmt.Println("[1].cur, cur.Next", cur.Val, (*cur.Next).Val, (*cur.Next) .Next)
	for cur != nil {

		pre, cur, cur.Next = cur, cur.Next, pre

		//fmt.Println("[1].pre, cur, cur.Next", pre.Val, (*pre.Next).Val, cur.Next )
		//fmt.Println("[2].pre, cur, cur.Next", pre.Val, (*pre.Next).Val, (*cur.Next) .Val, (*cur.Next).Next)
	}
	return pre
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre, cur *ListNode

	// head -> pre -> cur
	// head: 1-> 2-> 3-> nil pre: nil cur: nil

	fmt.Println("[0].",head, pre, cur)

	for head != nil {
		cur = head.Next
		// head: 1-> 2-> 3-> nil pre: nil cur: 2-> 3
		// head: 2->3 pre: 1->nil cur: 3-> nil
		// head: 3->nil pre: 2->1 cur: nil
		fmt.Println("[1].",head, pre, cur)
		head.Next = pre
		// head: 1-> nil pre: nil cur: 2-> 3
		// head: 2->1 pre: 1->nil cur: 3-> nil
		// head: 3->2->1 pre: 2->1 cur: nil
		fmt.Println("[2].",head, pre, cur)
		// head: 1-> nil pre: 1-> nil cur: 2-> 3
		// head: 2->1 pre: 2->1 cur: 3-> nil
		// head: 3->2->1 pre: nil cur: nil
		pre = head
		fmt.Println("[3].",head, pre, cur)
		head = cur
		// head: 2->3 pre: 1->nil cur: 2-> 3
		// head: 3->nil pre: 2->1 cur: 3-> nil
		// head: nil pre: 3->2->1 cur: nil
		fmt.Println("[4].",head, pre, cur)
	}

	return pre

}

/*
流程：
head: 1-> 2-> 3-> nil pre: nil cur: 2-> 3
head: 1-> 2-> 3-> nil  pre: nil cur: 2-> 3
head: 1-> nil  pre: nil cur: 2-> 3
head: 1->nil  pre: 1->nil cur: 2-> 3
head: 2->3  pre: 1->nil cur: 2-> 3
head: 2->3  pre: 1->nil cur: 3
head: 2->1-> nil  pre: 1->nil cur: 3
head: 2->1-> nil  pre: 2->1-> nil cur: 3
head: 3-> nil  pre: 2->1-> nil cur: nil
head: 3-> 2->1-> nil  pre: 2->1-> nil cur: nil
head: 3-> 2->1-> nil  pre: 3->2->1-> nil cur: nil
head: nil  pre: 3->2->1-> nil cur: nil
 */


func reverseList(head *ListNode) *ListNode {
	if head.Next == nil || head == nil {
		return head
	}
	var pre, cur *ListNode
	if head != nil {
		cur = head.Next
		head.Next = pre
		pre = head
		head = cur
	}
	return pre
}

func main() {

	cur := new(ListNode)
	cur.Val = 3
	cur.Next = nil

	pre := new(ListNode)
	pre.Val = 2
	pre.Next = cur

	head := new(ListNode)
	head.Val = 1
	head.Next = pre

	fmt.Println(head)
	fmt.Println(reverseList(head))
}
