package main

// This is the struct of the input root. Do not edit it.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	d := 0
	getLengthFromNode(root, &d)
	return d
}

func getLengthFromNode(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}

	lengthRight := getLengthFromNode(root.Right, maxDiameter)
	lengthLeft := getLengthFromNode(root.Left, maxDiameter)

	*maxDiameter = max(*maxDiameter, lengthLeft+lengthRight)
	return max(lengthLeft, lengthRight) + 1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
