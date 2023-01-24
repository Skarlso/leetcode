package main

func nextGreatestLetter(letters []byte, target byte) byte {
	// So the min should be the greatest in the list.
	// But that's fine because the last one would be the greatest
	// because it's a sorted list.
	min := letters[len(letters)-1]
	var found bool
	for _, v := range letters {
		if v > target && v <= min {
			found = true
			min = v
		}
	}
	if !found {
		return letters[0]
	}
	return min
}

// Since it's sorted we can do a modified binary search.
func nextGreatestLetterV2(letters []byte, target byte) byte {
	// return early if the target is out of bounds.
	// Sorted ascending.
	if target >= letters[len(letters)-1] || target < letters[0] {
		return letters[0]
	}

	low := 0
	high := len(letters) - 1
	for low <= high {
		mid := (low + high) / 2

		if target >= letters[mid] {
			low = mid + 1
		}
		if target < letters[mid] {
			high = mid - 1
		}
	}

	return letters[low]
}
