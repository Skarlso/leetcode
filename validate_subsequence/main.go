package main

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	var pop int
	pop, sequence = sequence[0], sequence[1:]

	for _, k := range array {
		if k == pop {
			if len(sequence) == 0 {
				return true
			}
			pop, sequence = sequence[0], sequence[1:]
		}
	}
	return false
}
