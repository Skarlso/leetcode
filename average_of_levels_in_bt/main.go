package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// if the depth is higher than the prev depth.
func averageOfLevels(root *TreeNode) []float64 {
	var (
		result  []float64
		current *TreeNode
	)
	queue := []*TreeNode{root}

	// sums := make(map[int]sum)
	// prevDepth := 0
	// sum := 0
	for len(queue) > 0 {
		l := len(queue)
		sum := 0
		for i := 0; i < l; i++ {
			current, queue = queue[0], queue[1:]
			sum += current.Val
			if current.Right != nil {
				queue = append(queue, current.Right)
			}

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
		}
		result = append(result, float64(sum)/float64(l))
		// l := len(queue)
		// sum := 0
		// current, queue = queue[0], queue[1:]
		// fmt.Println("depth: ", current.depth)
		// fmt.Println("=====")
		// if current.depth > prevDepth {

		// }

		// result = append(result, float64(current.sum))
		// if current.prev != nil && current.depth == current.prev.depth {
		// current.sum += current.node.Val
		// } else if current.prev != nil && current.depth > current.prev.depth {
		// 	result = append(result, float64(sum))
		// 	sum = 0
		// }
		// if _, ok := sums[current.depth]; !ok {
		// 	prevDepth := current.depth - 1
		// 	if v, ok := sums[prevDepth]; ok {
		// 		result = append(result, float64(v.sum)/float64(v.nodeCount))
		// 	}
		// }
		// s := sums[current.depth]
		// s.sum += current.node.Val
		// s.nodeCount++
		// sums[current.depth] = s

		// if current.node.Left != nil {
		// 	sum += current.node.Left.Val
		// 	queue = append(queue, &leaf{
		// 		depth: current.depth + 1,
		// 		node:  current.node.Left,
		// 	})
		// }

		// if current.node.Right != nil {
		// 	sum += current.node.Right.Val
		// 	queue = append(queue, &leaf{
		// 		depth: current.depth + 1,
		// 		node:  current.node.Right,
		// 	})
		// }

		// result = append(result, float64(sum)/float64(l))
	}

	return result
}
