package main

// I need some kind of cache for all the values?
type NumArray struct {
	nums []int
}

// Build a an array of sums from the values that are present in the array.
// After that we just have to recall each number by subtracting them.
func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.nums[right]
	}

	return this.nums[right] - this.nums[left-1]
}
