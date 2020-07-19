package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1->2->3->4->nil
// nil<-1 2->3->4->nil
func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode
	var after *ListNode
	cur := head

	for cur != nil {
		after = cur.Next // 2
		cur.Next = pre //

		pre = cur
		cur = after

		fmt.Println(cur, after)
	}

	fmt.Println("pre", pre)
	return pre


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

	fmt.Println(tmp)


}

