package main

// Definition for singly-linked list.
type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	newRoot := root
	queue := []*TreeNode{newRoot}
	var current *TreeNode

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		current.Left, current.Right = current.Right, current.Left

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return newRoot

}

func invertTreeRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	if root.Left != nil {
		root.Left = invertTreeRecursive(root.Left)
	}

	if root.Right != nil {
		root.Right = invertTreeRecursive(root.Right)
	}

	return root
}
