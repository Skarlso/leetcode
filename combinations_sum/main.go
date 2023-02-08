package main

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
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
	backtrack(path, result, candidates, index, target, currentTarget+candidates[index])
	path = path[:len(path)-1]

	// next decision where we don't include the candidate ( by moving the index ) to avoid duplicates and we also
	// don't change the total, because we didn't add in the candidate.
	backtrack(path, result, candidates, index+1, target, currentTarget)
}
