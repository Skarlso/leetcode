package main

func permute(nums []int) [][]int {
	var result [][]int
	// build up a counter, aka. frequency map to track duplications.
	counter := make(map[int]int)
	for _, v := range nums {
		counter[v]++
	}
	backtrack(&result, len(nums), counter, []int{})
	return result
}

func backtrack(result *[][]int, l int, counter map[int]int, path []int) {
	// we reached the number of items we want in a single path
	// add it and stop the recursion.
	if len(path) == l {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for k, v := range counter {
		// if we have an item we didn't encounter yet, we add it to a path.
		if v > 0 {
			// insert element in the path and decrease its frequency, meaning, we encountered the element.
			path = append(path, k)
			counter[k]--
			// do it again with the modified list. Anything that is 0 has already been encountered
			// so we will skip those. Anything that is above 0 will be added to the path until its full.
			backtrack(result, l, counter, path)
			// once a chain is completed, return and remove the item we added and increase its count again so
			// we can add it on the next loop.
			path = path[:len(path)-1]
			counter[k]++
		}
	}
}

// func permute(nums []int) [][]int {
// 	var result [][]int
// 	// pick a number and remove it from the list, then insert it into every position in the remaining list.
// 	sort.Ints(nums)
// 	backtrack(nums, &result, 0, map[int]bool{})
// 	return result
// }

// func backtrack(nums []int, result *[][]int, index int, seen map[int]bool) {
// 	// In Go we have to copy the slice or else we just change the underlying array
// 	// skewering the result.
// 	if len(nums) == index {
// 		temp := make([]int, len(nums))
// 		copy(temp, nums)
// 		*result = append(*result, temp)
// 	}

// 	for i := index; i < len(nums); i++ {
// 		if seen[index] {
// 			continue
// 		}
// 		if i > 0 && nums[i] == nums[i-1] && !seen[index-1] {
// 			continue
// 		}
// 		// add the current number into the slice
// 		nums[index], nums[i] = nums[i], nums[index]
// 		seen[index] = true
// 		// do the whole thing again but start from the next index
// 		backtrack(nums, result, index+1, seen)
// 		// remove what we did and do it again with the next item.
// 		nums[i], nums[index] = nums[index], nums[i]
// 		seen[index] = false
// 	}
// }
