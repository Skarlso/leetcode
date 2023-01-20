package main

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.
	queue := []*Node{n}
	seen := make(map[*Node]struct{})
	var current *Node

	for len(queue) > 0 {
		current, queue = queue[len(queue)-1], queue[:len(queue)-1]
		array = append(array, current.Name)

		// Go in reverse because we want the children from Left -> Right.
		for i := len(current.Children) - 1; i >= 0; i-- {
			if _, ok := seen[current.Children[i]]; !ok {
				queue = append(queue, current.Children[i])
				seen[current.Children[i]] = struct{}{}
			}
		}
	}

	return array
}
