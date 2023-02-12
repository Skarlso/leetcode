package main

type key struct {
	index, total int
}

func findTargetSumWays(nums []int, target int) int {
	dp := make(map[key]int)
	return backtrack(nums, 0, target, 0, dp)
}

func backtrack(candidates []int, index, target, total int, dp map[key]int) int {
	if index >= len(candidates) {
		if total == target {
			return 1
		}
		return 0
	}

	// another base case is if the pair already exists in cache, we just return the value.
	if v, ok := dp[key{index: index, total: total}]; ok {
		return v
	}

	dp[key{index: index, total: total}] = backtrack(candidates, index+1, target, total+candidates[index], dp) +
		backtrack(candidates, index+1, target, total-candidates[index], dp)

	return dp[key{index: index, total: total}]
}
