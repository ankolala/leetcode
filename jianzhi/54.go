package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	orderList := []int{}

	inOrder(root, &orderList)
	fmt.Println(orderList)
	return orderList[k-1]

}

func inOrder(root *TreeNode, orderList *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Right, orderList)
	*orderList = append(*orderList, root.Val)
	inOrder(root.Left, orderList)

}

func main() {
	left := &TreeNode{
		Val: 1,
		Left: nil,
		Right: nil,
	}

	lright := &TreeNode{
		Val: 2,
		Left: nil,
		Right: nil,
	}

	rright := &TreeNode{
		Val: 5,
		Left: nil,
		Right: nil,
	}

	right := &TreeNode{
		Val: 4,
		Left: lright,
		Right: rright,
	}

	top := &TreeNode{
		Val: 3,
		Left:left,
		Right:right,

	}
	fmt.Println(kthLargest(top, 1))
}

