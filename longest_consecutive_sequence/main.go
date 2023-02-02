package main

// apparently, it's faster to iterate over the map.
func longestConsecutive(nums []int) int {
	// first pass store the numbers in a map
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	// second pass, recursively find numbers of +1
	// this is just contains basically...
	max := 0
	for k := range m {
		if !m[k-1] {
			n := k
			for m[n] {
				n++
			}
			if n-k > max {
				max = n - k
			}
		}
	}

	return max
}
