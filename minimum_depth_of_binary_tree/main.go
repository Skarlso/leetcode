package main

// This is the struct of the input root. Do not edit it.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Leaf wraps a node item to track the depth with it. Normally, this would be tracked
// inside the tree itself and increase during addition. But now, we have to wrap it
// because we can't alter the original tree structure.
// I'm sure we could be doing this a bit differently by tracking the depth in a separate
// struct or queue or something. But this is a lot simpler to manage.
type leaf struct {
	prev  *leaf
	node  *TreeNode
	depth int
}

func minDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Write your code here.
	queue := []*leaf{{node: root, depth: 1}}
	var current *leaf

	for len(queue) > 0 {
		// Pop the last element.
		current, queue = queue[len(queue)-1], queue[:len(queue)-1]

		// Our Current Depth is the previous depth +1. Take care that this should not increase
		// by one for every node because that would be an aggregate. This is +1 per level. And
		// not per node.
		if current.prev != nil {
			current.depth = current.prev.depth + 1
		}

		if current.node.Left == nil && current.node.Right == nil {
			return current.depth
		}

		// Get the Right and Left node and add them to the queue.
		if current.node.Left != nil {
			queue = append(queue, &leaf{
				node: current.node.Left,
				prev: current,
			})
		}

		if current.node.Right != nil {
			queue = append(queue, &leaf{
				node: current.node.Right,
				prev: current,
			})
		}
	}

	return 0
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func minDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + minDepthRecursive(root.Right)
	} else if root.Right == nil {
		return 1 + minDepthRecursive(root.Left)
	}
	return 1 + min(minDepthRecursive(root.Left), minDepthRecursive(root.Right))
}
