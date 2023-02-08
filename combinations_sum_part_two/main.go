package main

import "sort"

// we faced something like this before where the number only must appear once.
// we used a counter or a frequency map for that.
func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backtrack([]int{}, &result, candidates, 0, target, 0)
	return result
}

func backtrack(path []int, result *[][]int, candidates []int, index, target, currentTarget int) {
	if currentTarget == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	if index >= len(candidates) || currentTarget > target {
		return
	}

	// first decision where we include the candidate.
	path = append(path, candidates[index])
	backtrack(path, result, candidates, index+1, target, currentTarget+candidates[index])
	path = path[:len(path)-1]

	// loop until we find a none duplicate
	for index+1 < len(candidates) && candidates[index] == candidates[index+1] {
		index++
	}

	// next decision where we don't include the candidate ( by moving the index ) to avoid duplicates and we also
	// don't change the total, because we didn't add in the candidate.
	backtrack(path, result, candidates, index+1, target, currentTarget)
}
