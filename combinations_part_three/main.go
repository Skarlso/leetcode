package main

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	counter := make(map[int]int)
	for _, n := range candidates {
		counter[n]++
	}
	backtrack([]int{}, &result, candidates, 0, n, 0, k, counter)
	return result
}

func backtrack(path []int, result *[][]int, candidates []int, index, target, currentTarget, height int, counter map[int]int) {
	if len(path) == height && currentTarget == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	if index >= len(candidates) || currentTarget > target {
		return
	}

	// first decision where we include the candidate.
	if counter[candidates[index]] > 0 {
		counter[candidates[index]]--
		path = append(path, candidates[index])
		backtrack(path, result, candidates, index, target, currentTarget+candidates[index], height, counter)
		path = path[:len(path)-1]
		counter[candidates[index]]++
	}
	// next decision where we don't include the candidate ( by moving the index ) to avoid duplicates and we also
	// don't change the total, because we didn't add in the candidate.
	backtrack(path, result, candidates, index+1, target, currentTarget, height, counter)
}
