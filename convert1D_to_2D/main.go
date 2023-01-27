package main

func construct2DArray(original []int, m int, n int) [][]int {
	l := len(original)
	if m*n != l {
		return [][]int{}
	}

	results := make([][]int, m)
	rows := 0
	for i := 0; i < m*n; i += n {
		row := original[i : i+n]
		results[rows] = row
		rows++
	}
	return results
}
