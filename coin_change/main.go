package main

func coinChange(coins []int, amount int) int {
	// initalize to the max value
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	// because it takes zero coins to get to zero
	dp[0] = 0

	for a := 1; a < amount+1; a++ {
		for _, c := range coins {
			if a-c >= 0 {
				// recurrence relationship
				dp[a] = min(dp[a], 1+dp[a-c])
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
