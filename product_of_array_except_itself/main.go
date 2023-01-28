package main

func productExceptSelf(nums []int) (ret []int) {
	for range nums {
		ret = append(ret, 1)
	}
	left := 1
	right := 1
	l := len(nums)
	for i := range nums {
		ret[i] *= left
		ret[l-1-i] *= right
		left *= nums[i]
		right *= nums[l-1-i]
	}
	return
}
