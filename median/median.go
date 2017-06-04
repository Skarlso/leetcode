package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	return median(nums1)
}

func median(arr []int) (med float64) {
	l := len(arr)
	if l < 1 {
		return 0.0
	}
	if l&1 == 1 {
		med = float64(arr[(l-1)/2])
	} else {
		med = float64((arr[(l-1)/2] + arr[l/2])) / 2
	}
	return
}

func main() {
	fmt.Println("Median: ", findMedianSortedArrays([]int{1, 3}, []int{2}))
}
