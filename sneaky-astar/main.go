package main

type point struct {
	x, y int
}

func (p *point) toSlice() []int {
	return []int{p.y, p.x}
}

// Normally, a* would use a heuristic approach with a Priority Queue.
// But I'm not implementing that in Go. :D
func AStarAlgorithm(startRow int, startCol int, endRow int, endCol int, graph [][]int) [][]int {
	start := point{x: startCol, y: startRow}
	goal := point{x: endCol, y: endRow}
	queue := []point{start}
	cameFrom := map[point]point{
		start: start,
	}
	var current point
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if current == goal {
			result := make([][]int, 0)

			for current != start {
				result = append(result, current.toSlice())
				current = cameFrom[current]
			}
			result = append(result, current.toSlice())
			for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
				result[i], result[j] = result[j], result[i]
			}
			return result
		}

		for _, next := range neighbours(current, graph) {
			if _, ok := cameFrom[next]; !ok {
				cameFrom[next] = current
				queue = append(queue, next)
			}
		}

	}
	return [][]int{}
}

var directions = []point{
	{x: 0, y: -1}, // up
	{x: 0, y: 1},  // down
	{x: 1, y: 0},  // left
	{x: -1, y: 0}, // right
}

func neighbours(p point, graph [][]int) []point {
	var result []point
	for _, d := range directions {
		current := point{x: p.x + d.x, y: p.y + d.y}
		if current.x >= 0 && current.y >= 0 && current.y < len(graph) && current.x < len(graph[current.y]) {
			if graph[current.y][current.x] != 1 {
				result = append(result, current)
			}
		}
	}

	return result
}
