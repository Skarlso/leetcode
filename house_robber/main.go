package main

func rob(nums []int) int {
	rob1, rob2 := 0, 0

	// [rob1, rob2, n, n+1]
	for _, n := range nums {
		rob1, rob2 = rob2, max(n+rob1, rob2)
	}
	return rob2
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
