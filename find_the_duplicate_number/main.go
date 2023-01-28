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

// divide and conquer
// do the thing where we go left from right, but constantly divide up the list.
func findDuplicate(nums []int) int {
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]&result == result {
			return nums[i]
		}

		result |= nums[i]
	}

	return -1
}

func findDuplicateMark(nums []int) int {
	l := len(nums)

	for _, v := range nums {
		index := abs(v)

		if nums[index] < 0 {
			return index
		}

		nums[index] = -nums[index]
	}

	return l
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// fast slow pointer...
public int findDuplicate_fastSlow(int[] nums) {
	int slow = 0;
	int fast = 0;
	do {
		slow = nums[slow];
		fast = nums[nums[fast]];
	} while (slow != fast);

	slow = 0;
	while (slow != fast) {
		slow = nums[slow];
		fast = nums[fast];
	}

	return slow;
}