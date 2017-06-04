package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := append(nums1, nums2...)
	sort.Ints(arr)
	med := median(arr)
	return med
}

func median(arr []int) (med float64) {
	if len(arr) < 1 {
		return 0.0
	}
	if len(arr)&1 == 1 {
		med = float64(arr[(len(arr)-1)/2])
	} else {
		med = float64((arr[(len(arr)-1)/2] + arr[(len(arr)/2)])) / 2
	}
	return
}

func main() {
	fmt.Println("Median: ", findMedianSortedArrays([]int{1, 3}, []int{2}))
}
