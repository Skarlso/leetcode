package main

func combine(n int, k int) [][]int {
	var result [][]int
	backtrack([]int{}, &result, 1, n, k)
	return result
}

func backtrack(path []int, result *[][]int, start, end int, hight int) {
	if len(path) == hight {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for i := start; i <= end; i++ {
		// add the current number into the slice
		path = append(path, i)
		// choose the next item. In order to avoid duplicates, we will choose the next item + 1 which is greater than
		// the item before it. It can only always increase and not go back.
		backtrack(path, result, i+1, end, hight)
		// remove what we did and do it again with the next item.
		path = path[:len(path)-1]
	}
}
