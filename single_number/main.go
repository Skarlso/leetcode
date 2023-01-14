package main

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

// You must implement a solution with a linear runtime complexity and use only constant extra space.
// An obvious solution is using a HashMap.
// The right solution is using bit magic.
// Another option is to sort the list and then compare them. -> That is also accepted.
func singleNumber(nums []int) int {
	var n int
	for _, i := range nums {
		n ^= i
	}
	return n
}
