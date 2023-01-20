package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	sums := []int{}
	queue := []*BinaryTree{root}
	cameFrom := make(map[*BinaryTree]*BinaryTree)
	cameFrom[root] = nil
	var current *BinaryTree

	for len(queue) > 0 {
		current, queue = queue[len(queue)-1], queue[:len(queue)-1]

		if current.Left == nil && current.Right == nil {
			sum := 0
			curr := current
			for curr != root {
				sum += curr.Value
				curr = cameFrom[curr]
			}
			sum += curr.Value
			sums = append(sums, sum)
		}

		for _, next := range []*BinaryTree{current.Right, current.Left} {
			if _, ok := cameFrom[next]; next != nil && !ok {
				cameFrom[next] = current
				queue = append(queue, next)
			}
		}
	}

	return sums
}

type leaf struct {
	node *BinaryTree
	sum  int
}

func BranchSumsV2(root *BinaryTree) []int {
	// Write your code here.
	sums := []int{}
	queue := []*leaf{{node: root, sum: 0}}
	var current *leaf

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if current.node.Left == nil && current.node.Right == nil {
			sums = append(sums, current.sum+current.node.Value)
		}

		if current.node.Left != nil {
			queue = append(queue, &leaf{
				node: current.node.Left,
				sum:  current.sum + current.node.Value,
			})
		}

		if current.node.Right != nil {
			queue = append(queue, &leaf{
				node: current.node.Right,
				sum:  current.sum + current.node.Value,
			})
		}
	}

	return sums
}
