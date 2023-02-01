package main

type point struct {
	x, y int
}

var (
	directions = []point{
		{x: -1, y: 0},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: 0, y: -1},
	}
)

// Recursive
func exist(board [][]byte, word string) bool {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] == word[0] {
				if search(board, word, point{x: x, y: y}) {
					return true
				}
			}
		}
	}
	return false
}

func search(board [][]byte, word string, p point) bool {
	if word == "" {
		return true
	}

	if p.x < 0 || p.y < 0 || p.y >= len(board) || p.x >= len(board[p.y]) || board[p.y][p.x] != word[0] {
		return false
	}

	prev := board[p.y][p.x]
	board[p.y][p.x] = '.' // mark as visited

	if search(board, word[1:], point{x: p.x + 1, y: p.y}) {
		return true
	}
	if search(board, word[1:], point{x: p.x - 1, y: p.y}) {
		return true
	}
	if search(board, word[1:], point{x: p.x, y: p.y + 1}) {
		return true
	}

	if search(board, word[1:], point{x: p.x, y: p.y - 1}) {
		return true
	}

	board[p.y][p.x] = prev // put back and try again
	return false
}

// func neighbor(p point, board [][]byte, char byte) []point {
// 	var result []point

// 	return result
// }

// func exist(board [][]byte, word string) bool {
// 	if word == "" {
// 		return false
// 	}

// 	first := word[0]

// 	startingPositions := make([]point, 0)
// 	for y := 0; y < len(board); y++ {
// 		for x := 0; x < len(board[y]); x++ {
// 			if board[y][x] == first {
// 				startingPositions = append(startingPositions, point{x: x, y: y})
// 			}
// 		}
// 	}

// 	var start point
// 	for len(startingPositions) > 0 {
// 		start, startingPositions = startingPositions[0], startingPositions[1:]

// 		queue := []point{start}
// 		seen := make(map[point]bool)
// 		var current point
// 		// this is the point where we should be backtracking.. but I wonder if this would work iteratively.
// 		for len(queue) > 0 {
// 			current, queue = queue[0], queue[1:]

// 			if board[current.y][current.y] == word[len(word)-1] {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }
