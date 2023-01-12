package main

// Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the
// range [1, n] that do not appear in nums.
func findDisappearedNumbers(nums []int) []int {
	// so this is find missing on steroids with multiple possible missing entries.
	// I can do something like, loop, put in dictionary, than loop from 1,n and
	// if the map doesn't have that number, put it into output.
	// this isn't that nice though.
	m := make(map[int]struct{})
	// or... sort it and if it doesn't equal with the index + 1, put the index into the output.
	// [1, 3, 5, 6]
	// index: 0, 1, 2, 3 -> instead of index, use a number that is a follower.
	// what are my outputs?
	// from the previous number to the current number.
	// 1-3, and then 3-5.
	// this requires a sorted list.
	var result []int
	// sort.Ints(nums) // this is a quicksort.
	// offset := 1 // this doesn't work because there could be duplicates
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] != offset {
	// 		for j := nums[i-1]; j <= nums[i]; j++ {
	// 			result = append(result, j)
	// 		}
	// 		offset = nums[i]
	// 	} else {
	// 		offset++
	// 	}
	// }

	for _, i := range nums {
		m[i] = struct{}{}
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			result = append(result, i)
		}
	}
	return result
}

func findDisappearedNumbersV2(nums []int) []int {
	for _, k := range nums {
		index := abs(k) - 1
		nums[index] = -abs(nums[index])
	}

	var result []int
	for i, num := range nums {
		if num > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
