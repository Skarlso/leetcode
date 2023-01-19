package main

import "sort"

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)

	sum := 0
	for _, coin := range coins {
		if sum+1 < coin {
			break
		}
		sum += coin
	}
	return sum + 1
}
