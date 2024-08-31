package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	j := 2
	for i := 2; i < n; i++ {
		if nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
