package main

type ListNode struct {
	Val     int
	Next    *ListNode
}


func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head

	for k>0 && fast != nil {
		k --
		fast = fast.Next
		if fast == nil {

			return slow
		}

	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow


}

func main() {

	// step = k-1

	head := &ListNode{
		Val : 1,
		Next: nil,
	}
	getKthFromEnd(head,1)
}

