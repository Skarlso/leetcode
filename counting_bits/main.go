package main

// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of
// 1's in the binary representation of i.
func countBitsv1(n int) []int {
	var result []int
	for i := 0; i <= n; i++ {
		num := i
		count := 0
		for num > 0 {
			count += num & 1
			num >>= 1
		}
		result = append(result, count)
	}
	return result
}

func countBits(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		result[i] = result[i/2] + i%2
	}
	return result
}
