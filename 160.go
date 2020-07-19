package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

// 相交链表

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	slow, fast := headA, headB

	for slow != fast {
		if (slow != nil) {
			slow = slow.Next
		} else {
			slow = headB
		}

		if (fast != nil) {
			fast = fast.Next
		} else {
			fast = headA
		}
	}
		return slow
}

func main() {
	headA := &ListNode{
	}
	headB := &ListNode{

	}
	fmt.Println(getIntersectionNode(headA, headB))
}
