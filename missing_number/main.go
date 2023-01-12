package main

// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is
// missing from the array.
func missingNumber(nums []int) int {
	// add them together, calculate the total sum from the numbers then subtract the actual sum
	// and what remains should be to number that is missing.
	sum := 0
	for _, i := range nums {
		sum += i
	}
	nnums := len(nums)
	// Arithmetic sum of list of numbers
	actualSum := (nnums * (nnums + 1)) / 2
	return actualSum - sum
}
