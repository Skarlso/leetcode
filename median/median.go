package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	l := len(nums1)
	if l < 1 {
		return 0.0
	}
	if l&1 == 1 {
		return float64(nums1[(l-1)/2])
	}
	return float64((nums1[(l-1)/2] + nums1[l/2])) / 2
}

func main() {
	fmt.Println("Median: ", findMedianSortedArrays([]int{1, 3}, []int{2}))
}
