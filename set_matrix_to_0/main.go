package main

type point struct {
	x, y int
}

type Range struct {
	from, to point
}

func setZeroes(matrix [][]int) {
	// use these two mark the rows and columns that need to be 0-ed.
	rowRange := make(map[Range]struct{})
	colRange := make(map[Range]struct{})

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == 0 {
				rowRange[Range{
					from: point{x: 0, y: y},
					to:   point{x: len(matrix[y]), y: y},
				}] = struct{}{}
				colRange[Range{
					from: point{x: x, y: 0},
					to:   point{x: x, y: len(matrix)},
				}] = struct{}{}
			}
		}
	}

	for k := range rowRange {
		for x := k.from.x; x < k.to.x; x++ {
			matrix[k.from.y][x] = 0
		}
	}

	for k := range colRange {
		for y := k.from.y; y < k.to.y; y++ {
			matrix[y][k.from.x] = 0
		}
	}
}

func setZeroConstantMemory(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	var (
		frow bool
		fcol bool
	)

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == 0 {
				if row == 0 {
					frow = true
				}
				if col == 0 {
					fcol = true
				}
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if matrix[0][col] == 0 || matrix[row][0] == 0 {
				matrix[row][col] = 0
			}
		}
	}

	if frow {
		for col := 0; col < n; col++ {
			matrix[0][col] = 0
		}
	}
	if fcol {
		for row := 0; row < m; row++ {
			matrix[row][0] = 0
		}
	}
}
