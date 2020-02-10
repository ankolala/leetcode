package main

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

func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		} else {
			hash[head] = true
		}
		head = head.Next
	}
	return false
}

func main() {

}