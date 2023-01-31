package main

func findDuplicates(nums []int) []int {
	var res []int

	for _, n := range nums {
		if nums[abs(n)-1] < 0 {
			res = append(res, abs(n))
		} else {
			nums[abs(n)-1] *= -1
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
