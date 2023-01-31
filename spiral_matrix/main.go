package main

func spiralOrder(matrix [][]int) []int {
	var result []int

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return result
	}

	var (
		n         = len(matrix)
		m         = len(matrix[0])
		top, left int
		bottom    = n - 1
		right     = m - 1
	)

	for len(result) < n*m {
		// add top
		for j := left; j <= right && len(result) < n*m; j++ {
			result = append(result, matrix[top][j])
		}

		// add right
		for j := top + 1; j <= bottom-1 && len(result) < n*m; j++ {
			result = append(result, matrix[j][right])
		}

		// add bottom
		for j := right; j >= left && len(result) < n*m; j-- {
			result = append(result, matrix[bottom][j])
		}

		// add left
		for j := bottom - 1; j >= top+1 && len(result) < n*m; j-- {
			result = append(result, matrix[j][left])
		}

		// adjust the barriers.
		left++
		right--
		top++
		bottom--
	}

	return result
}
