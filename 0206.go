package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}

// 判断是否为回文链表

// 快慢指针：找到链表中点, 反转链表，依次作比较
// 额外开辟数组空间

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	left, right := 0, len(result) - 1

	for left < right {
		if result[left] != result[right]{
			return false
		}
		left ++
		right --

	}
	return true
}

func main() {
	// 2->1 false
	// 2->1->2->5 true
	// 2->4->3->1->3->4->2 true
	// 3->1->1->3

	fourth := &ListNode{
		Val: 3,
		Next: nil,

	}

	third := &ListNode{
		Val: 1,
		Next: fourth,

	}
	second := &ListNode{
		Val: 1,
		Next: third,

	}

	head := &ListNode{
		Val: 3,
		Next: second,
	}

	fmt.Println(isPalindrome(head))
}
