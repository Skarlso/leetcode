package main

// This is the struct of the input root. Do not edit it.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type leaf struct {
	node *TreeNode
	sum  int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// Write your code here.
	queue := []*leaf{{node: root, sum: 0}}
	var current *leaf

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if current.node.Left == nil && current.node.Right == nil {
			if current.sum+current.node.Val == targetSum {
				return true
			}
		}

		if current.node.Left != nil {
			queue = append(queue, &leaf{
				node: current.node.Left,
				sum:  current.sum + current.node.Val,
			})
		}

		if current.node.Right != nil {
			queue = append(queue, &leaf{
				node: current.node.Right,
				sum:  current.sum + current.node.Val,
			})
		}
	}

	return false
}

// Hey, but at least I came up with this one all by my self. :)
func hasPathSumRec(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return helper(root, targetSum, 0)
}

func helper(node *TreeNode, targetSum, currentSum int) bool {
	if node.Right == nil && node.Left == nil {
		return targetSum == currentSum+node.Val
	}
	right := false
	left := false
	if node.Right != nil {
		right = helper(node.Right, targetSum, currentSum+node.Val)
	}
	if node.Left != nil {
		left = helper(node.Left, targetSum, currentSum+node.Val)
	}

	return right || left
}

// This is easier on the eyes... Remember that you can also SUBTRACT dang it!
// func hasPathSumRec(root *TreeNode, sum int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	if root.Left == nil && root.Right == nil && sum-root.Val == 0 {
// 		return true
// 	}

// 	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
// }
