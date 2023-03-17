package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = 1
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	m := dp[0]
	for _, v := range dp {
		m = max(m, v)
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
