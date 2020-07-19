package main

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {

	// root为空时返回深度为0
	if root == nil {
		return 0
	}

	// 最大深度初始值
	m := 0

	depth(root, &m)

	return m
}


func depth(root *TreeNode, max *int) int {

	if root == nil {
		return 0
	}

	left := depth(root.Left, max)
	right := depth(root.Right, max)

	*max = maxInt(left+right, *max)

	return maxInt(left,right) + 1
}


func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

	diameterOfBinaryTree(&TreeNode{Val:1,Left:nil,Right:nil})
}