package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	reachedLength := 0
	var (
		maxAverage   float64 = math.MinInt64
		runningTotal float64
	)

	lastAddedIndex := 0

	for _, n := range nums {
		runningTotal += float64(n)
		reachedLength++

		if reachedLength == k {
			if runningTotal/float64(k) > maxAverage {
				maxAverage = runningTotal / float64(k)
			}

			runningTotal -= float64(nums[lastAddedIndex])

			lastAddedIndex++
			reachedLength--
		}
	}

	return maxAverage
}
