package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1->2->3->4
// 2->1->4->3
func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	if head.Next.Next == nil {
		// 1->2
		next := head.Next // next = 2
		head.Next = nil
		tmp := head // tmp = 1
		next.Next = tmp
		return next
	}
	// 1->2->3->4
	// 2->1->4->3
	next := head.Next // 2->3->4
	head.Next = swapPairs(next.Next) // 1->3->4
	tmp := head
	next.Next = tmp

	return next

}

func main(){

	nnn := new(ListNode)
	nnn.Val = 4
	nnn.Next = nil

	nextnext := new(ListNode)
	nextnext.Val = 3
	nextnext.Next = nnn

	next := new(ListNode)
	next.Val = 2
	next.Next = nextnext

	head := new(ListNode)
	head.Next = next
	head.Val = 1

	tmp := swapPairs(head)
	fmt.Println(tmp.Val, tmp.Next.Val, tmp.Next.Next.Val)


}

