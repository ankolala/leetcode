package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}

	tag1 := (left.Val == right.Val)
	tag2 := isMirror(left.Left, right.Right)
	tag3 := isMirror(left.Right, right.Left)

	return tag1&&tag2&&tag3
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func main(){

	cur := new(TreeNode)
	cur.Val = 2
	cur.Right = nil
	cur.Left = nil

	pre := new(TreeNode)
	pre.Val = 2
	pre.Left = nil
	pre.Right = nil

	head := new(TreeNode)
	head.Val = 1
	head.Left = cur
	head.Right = pre

	fmt.Println(isSymmetric(head))
}
