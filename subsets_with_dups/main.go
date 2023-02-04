package main

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	// This is a similar pattern to all backtracking based solutions.
	sort.Ints(nums)
	backtrack(nums, []int{}, &result, 0)
	return result
}

func backtrack(nums []int, output []int, result *[][]int, index int) {
	// In Go we have to copy the slice or else we just change the underlying array
	// skewering the result.
	temp := make([]int, len(output))
	copy(temp, output)
	// use a pointer to get to the answer, or return it back modified.
	// using a pointer is more efficient and reads easier.
	*result = append(*result, temp)
	for i := index; i < len(nums); i++ {
		// use a pointer to get to the answer, or return it back modified.
		// using a pointer is more efficient and reads easier.
		if i != index && nums[i] == nums[i-1] {
			continue
		}
		// add the current number into the slice
		output = append(output, nums[i])
		// do the whole thing again but start from the next index
		backtrack(nums, output, result, i+1)
		// remove what we did and do it again with the next item.
		output = output[:len(output)-1]
	}
}
