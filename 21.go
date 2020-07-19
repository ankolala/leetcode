package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
// 1->2->3
// 2->3->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1 == nil && l2 == nil {
		return nil
	}


	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l2.Next, l1)
	return l2
}

func main() {

	l0 := &ListNode{
		Val : 3,
		Next: nil,
	}

	l1 := &ListNode{
		Val : 1,
		Next: l0,
	}
	l2 := &ListNode{
		Val: 2,
		Next: nil,
	}

	//[1]
	//[2]
	//[1,2]
	node := mergeTwoLists(l1, l2)
	fmt.Println(node, node.Next, node.Next.Next)
}
