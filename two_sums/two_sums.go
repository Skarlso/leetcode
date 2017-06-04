package main

import "fmt"

func main() {
	fmt.Println("Bla:", twoSum2([]int{-1, -1, -2, -3, -4, -5}, -8))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	myMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		myMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if val, ok := myMap[complement]; ok {
			if val != i {
				return []int{i, val}
			}
		}
	}
	return []int{}
}
