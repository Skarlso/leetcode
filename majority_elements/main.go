package main

import (
	"math"
	"sort"
)

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	max := math.MinInt
	maxNum := 0
	currentStreak := 1
	for i := 0; i < len(nums); i++ {
		if i+1 == len(nums) {
			if nums[i] == nums[i-1] {
				currentStreak++
				if currentStreak > max && currentStreak > len(nums)/2 {
					max = currentStreak
					maxNum = nums[i]
				}
			}
		} else if nums[i] == nums[i+1] {
			currentStreak++
			if currentStreak > max && currentStreak > len(nums)/2 {
				max = currentStreak
				maxNum = nums[i]
			}
		} else {
			currentStreak = 1
		}
	}
	return maxNum
}

func majorityElementV2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	max := math.MinInt
	maxNum := 0
	currentStreak := 1
	for i, j := 0, 1; i < len(nums); i, j = i+1, j+1 {
		// if our second pointer gets to the end we check the previous
		// elements instead.
		if j == len(nums) {
			j -= 2
		}
		if nums[i] == nums[j] {
			currentStreak++
			if currentStreak > max && currentStreak > len(nums)/2 {
				max = currentStreak
				maxNum = nums[i]
			}
		} else {
			currentStreak = 1
		}
	}
	return maxNum
}

func majorityElementV3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxNum := nums[0]
	currentStreak := 1
	for i := 1; i < len(nums); i++ {
		if currentStreak == 0 {
			currentStreak++
			maxNum = nums[i]
		} else if maxNum == nums[i] {
			currentStreak++
		} else {
			currentStreak--
		}
	}
	return maxNum
}
