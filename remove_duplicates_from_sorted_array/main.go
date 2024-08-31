package main

func removeDuplicates(nums []int) int {
	i, j := 0, 0
	ret := 1
	for j < len(nums) {
		if nums[j] > nums[i] {
			nums[i+1], nums[j] = nums[j], nums[i+1]
			i++
			ret++
		}

		j++

	}

	return ret

}
