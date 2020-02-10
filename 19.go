package main

import "fmt"

/*
1->2->3->4->5, n=2

1->2->3->5

 */

type ListNode struct {
	Val int
	Next *ListNode
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 定义哑结点
	dummy := new(ListNode)
	dummy.Next = head
	length := 0
	cur := head
	for cur != nil {
		length ++
		cur = cur.Next
	}

	// 剔除idx的正向位置
	length = length - n
	fmt.Println(length)
	cur = dummy

	fmt.Println(cur)
	for length>0 {
		length --
		cur = cur.Next
	}
	fmt.Println(cur.Val, cur.Next.Val)
	cur.Next = cur.Next.Next
	return dummy.Next

}




func main() {

	head3 := new(ListNode)
	head3.Val = 5
	head3.Next = nil

	pre2 := new(ListNode)
	pre2.Val = 4
	pre2.Next = head3

	head2 := new(ListNode)
	head2.Val = 3
	head2.Next = pre2

	pre1 := new(ListNode)
	pre1.Val = 2
	pre1.Next = head2

	head1:= new(ListNode)
	head1.Val = 1
	head1.Next = pre1

	result := removeNthFromEnd(head1,4)

	fmt.Println(result, result.Next, result.Next.Next, result.Next.Next.Next)
}