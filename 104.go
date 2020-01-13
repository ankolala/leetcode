package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var depth = 0

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	right := maxDepth(root.Right)
	left := maxDepth(root.Left)
	depth = int(math.Max(float64(left),float64(right)))+1
	return depth
}

func main() {

	cur := new(TreeNode)
	cur.Val = 3
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

	fmt.Println(maxDepth(head))
}
