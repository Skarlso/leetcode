package main

// Definition for singly-linked list.
type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func isMatch(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}

	return isMatch(root.Left, subRoot.Left) && isMatch(root.Right, subRoot.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	// Check all nodes from root node.
	if isMatch(root, subRoot) {
		return true
	}
	// Check all branches for possible matches. From all branches do a Matching search
	// to see if all branches match up with the subRoot FROM subRoot.
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
