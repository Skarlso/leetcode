package main

// This is the struct of the input root. Do not edit it.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	// Write your code here.
	queue1 := []*TreeNode{p}
	queue2 := []*TreeNode{q}
	var current1, current2 *TreeNode

	for {
		if len(queue1) == 0 && len(queue2) == 0 {
			return true
		}

		if len(queue1) == 0 && len(queue2) != 0 || len(queue1) != 0 && len(queue2) == 0 {
			break
		}

		current1, queue1 = queue1[0], queue1[1:]
		current2, queue2 = queue2[0], queue2[1:]

		if current1 == nil || current2 == nil {
			return false
		}

		if current1.Left != nil && current2.Left == nil || current1.Left == nil && current2.Left != nil || current1.Right != nil && current2.Right == nil || current1.Right == nil && current2.Right != nil || current1.Val != current2.Val {
			return false
		}

		for _, next := range []*TreeNode{current1.Left, current1.Right} {
			if next != nil {
				queue1 = append(queue1, next)
			}
		}

		for _, next := range []*TreeNode{current2.Left, current2.Right} {
			if next != nil {
				queue2 = append(queue2, next)
			}
		}
	}

	return false
}

func isSameTreeRec(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return isSameTreeRec(p.Right, q.Right) && isSameTreeRec(p.Left, q.Left)
}
