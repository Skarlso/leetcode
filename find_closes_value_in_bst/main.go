package main

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// FindClosestValueV2 this one uses O(1) memory.
func (tree *BST) FindClosestValueV2(target int) int {
	// Write your code here.
	min := tree.Value
	for tree != nil {
		if abs(min-target) > abs(tree.Value-target) {
			min = tree.Value
		}

		if target < tree.Value {
			tree = tree.Left
		} else {
			tree = tree.Right
		}
	}

	return min
}
func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	queue := []*BST{tree}
	seen := make(map[*BST]struct{})
	minDiff := math.MaxInt
	minValue := tree.Value
	var current *BST
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		// if
		if abs(current.Value-target) < minDiff {
			minValue = current.Value
			minDiff = abs(current.Value - target)
		}

		for _, next := range []*BST{current.Left, current.Right} {
			if next != nil {
				if _, ok := seen[next]; !ok {
					queue = append(queue, next)
					seen[next] = struct{}{}
				}
			}
		}
	}
	return minValue
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
