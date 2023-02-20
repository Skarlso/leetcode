package main

func maxProduct(nums []int) int {
	res := max(nums...)
	curMax, curMin := 1, 1

	for _, n := range nums {
		curMax, curMin = max(curMax*n, curMin*n, n), min(curMax*n, curMin*n, n)
		res = max(res, curMax)
	}

	return res
}

func max(n ...int) int {
	m := n[0]
	for _, i := range n {
		if i > m {
			m = i
		}
	}

	return m
}

func min(n ...int) int {
	m := n[0]
	for _, i := range n {
		if i < m {
			m = i
		}
	}

	return m
}
