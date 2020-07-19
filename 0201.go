package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 定义hashMap保证链表元素唯一
	hashMap := map[int]int{}
	pre, cur := head, head.Next
	hashMap[pre.Val] = 1
	for cur != nil {
		// 不存在，赋值
		if _, ok := hashMap[cur.Val]; !ok{
			hashMap[cur.Val] = 1
			pre = pre.Next

		} else {
			pre.Next = cur.Next

		}
		cur = cur.Next

	}

	return head
}



func main() {

	data := []int{1, 1, 1, 1, 2}

	head := createLinkList(data)

	fmt.Println(printLinkList(head))


	//fmt.Println(removeDuplicateNodes(head))

}


func createLinkList(data []int) *ListNode {

	head := &ListNode{Val: data[0], Next: nil}
	cur := head

	for i := 1; i < len(data); i ++ {
		tmpNode := &ListNode{Val: data[i], Next: nil}
		cur.Next = tmpNode
		cur = cur.Next
	}

	return head
}

func printLinkList(head *ListNode) string {
	res := ""
	for head != nil {
		res += fmt.Sprintf("%d -> ", head.Val)
		head = head.Next
	}

	res += "nil"

	return res

}