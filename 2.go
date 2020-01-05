package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//1. l1, l2 为空
	if l1 == nil && l2 == nil {
		return nil
	}

	//2. l1为空
	if l1 == nil {
		return l2
	}

	//3. l2为空
	if l2 == nil {
		return l1
	}
	//4. l1 与 l2 不为空
	sum := l1.Val + l2.Val
	NextNode := addTwoNumbers(l1.Next, l2.Next)

	if sum < 10 {
		return &ListNode{
			Val : sum,
			Next: NextNode,
		}
	} else {
		tmpNode := &ListNode{
			Val : 1,
			Next : nil,
		}
		return &ListNode{
			Val : sum-10,
			Next : addTwoNumbers(NextNode, tmpNode),
		}
	}

}

func main() {
	/*cur2 := new(ListNode)
	cur2.Val = 3
	cur2.Next = nil
	*/
	pre2 := new(ListNode)
	pre2.Val = 7
	pre2.Next = nil


	head2 := new(ListNode)
	head2.Val = 3
	head2.Next = pre2
	//head2.Next = nil

	/*cur1 := new(ListNode)
	cur1.Val = 4
	cur1.Next = nil
	*/
	pre1 := new(ListNode)
	pre1.Val = 2
	//pre1.Next = cur1
	pre1.Next = nil

	head1:= new(ListNode)
	head1.Val = 9
	head1.Next = pre1
	//head1.Next = nil

	head_new := addTwoNumbers(head1, head2)
	fmt.Println(head_new,head_new.Next)
	/*
	[2,4,3] -> [4,3,2]
	[5,6,4] -> [6,5,4]
	[7,10,7]
	 */
	//fmt.Println(head_new, head_new.Next, head_new.Next.Next)
	//fmt.Println(head_new,head_new.Next, head_new.Next.Next)
}
