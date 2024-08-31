package main

// removeElement uses two pointers one from going the opposite until finding
// a value that is good and stops there, than continues with the other
// pointer.
func removeElement(nums []int, val int) int {
	// handle special case
	l := len(nums) - 1
	if l == 0 {
		if nums[0] == val {

			return 0
		}

		return 1
	}

	i, j := 0, l
	ret := l + 1
	for i <= j {
		if nums[j] == val {
			j--

			// continue the end pointer until we find a good number to
			// change with.
			// also, increase ret because we just found a number
			ret--
			continue
		}

		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--

			ret--
		}

		i++
	}

	return ret
}
