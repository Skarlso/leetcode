package main

func permute(nums []int) [][]int {
	var result [][]int
	// pick a number and remove it from the list, then insert it into every position in the remaining list.
	backtrack(nums, &result, 0)
	return result
}

func backtrack(nums []int, result *[][]int, index int) {
	// In Go we have to copy the slice or else we just change the underlying array
	// skewering the result.
	if len(nums) == index {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
	}

	for i := index; i < len(nums); i++ {
		// add the current number into the slice
		nums[index], nums[i] = nums[i], nums[index]
		// do the whole thing again but start from the next index
		backtrack(nums, result, i+1)
		// remove what we did and do it again with the next item.
		nums[i], nums[index] = nums[index], nums[i]
	}
}
