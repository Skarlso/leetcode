package main

func construct2DArray(original []int, m int, n int) [][]int {
	var results [][]int

	l := len(original)
	if m*n != l {
		return [][]int{}
	}

	for i := 0; i < m*n; i += n {
		row := original[i : i+n]
		results = append(results, row)
	}

	if len(results) != m {
		return [][]int{}
	}
	return results
}
